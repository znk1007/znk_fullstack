package core

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

// 有效荷载编码
type payloadEncoder struct {
	supportBinary bool
	writeM        writeManager
	dt            pbs.DataType
	pt            pbs.PacketType
	header        bytes.Buffer
	cache         bytes.Buffer
	base64Writer  io.WriteCloser
	rawWriter     io.Writer
}

func (pe *payloadEncoder) Noop() []byte {
	if pe.supportBinary {
		return []byte{0x00, 0x01, 0xff, '6'}
	}
	return []byte("1:6")
}

func (pe *payloadEncoder) NextWriter(dt pbs.DataType, pt pbs.PacketType) (io.WriteCloser, error) {
	w, err := pe.writeM.getWriter()
	if err != nil {
		return nil, err
	}
	pe.rawWriter = w
	pe.dt = dt
	pe.pt = pt
	pe.cache.Reset()
	if !pe.supportBinary && dt == pbs.DataType_binary {
		pe.base64Writer = base64.NewEncoder(base64.StdEncoding, &pe.cache)
	} else {
		pe.base64Writer = nil
	}
	return pe, nil
}

func (pe *payloadEncoder) Write(p []byte) (int, error) {
	if pe.base64Writer != nil {
		return pe.base64Writer.Write(p)
	}
	return pe.cache.Write(p)
}

func (pe *payloadEncoder) Close() error {
	if pe.base64Writer != nil {
		pe.base64Writer.Close()
	}
	var writeHeader func() error
	if pe.supportBinary {
		writeHeader = pe.writeBinaryHeader
	} else {
		if pe.dt == pbs.DataType_binary {
			writeHeader = pe.writeBase64Header
		} else {
			writeHeader = pe.writeTextHeader
		}
	}
	pe.header.Reset()
	err := writeHeader()
	if err == nil {
		_, err = pe.header.WriteTo(pe.rawWriter)
	}
	if err == nil {
		_, err = pe.cache.WriteTo(pe.rawWriter)
	}
	if werr := pe.writeM.addWriter(err); werr != nil {
		return werr
	}
	return err
}

func (pe *payloadEncoder) writeTextHeader() error {
	l := int64(pe.cache.Len() + 1)
	err := writeTextLen(l, &pe.header)
	if err == nil {
		err = pe.header.WriteByte(packetTypeToASCIIByte(pe.pt))
	}
	return err
}

func (pe *payloadEncoder) writeBase64Header() error {
	l := int64(pe.cache.Len() + 2)
	err := writeTextLen(l, &pe.header)
	if err == nil {
		err = pe.header.WriteByte('b')
	}
	if err == nil {
		err = pe.header.WriteByte(packetTypeToASCIIByte(pe.pt))
	}
	return err
}

func (pe *payloadEncoder) writeBinaryHeader() error {
	l := int64(pe.cache.Len() + 1)
	b := packetTypeToASCIIByte(pe.pt)
	if pe.dt == pbs.DataType_binary {
		b = packetTypeToBinaryByte(pe.pt)
	}

	err := pe.header.WriteByte(dataTypeToByte(pe.dt))
	if err == nil {
		err = writeBinaryLen(l, &pe.header)
	}
	if err == nil {
		err = pe.header.WriteByte(b)
	}
	return err
}

type payloadDecoder struct {
	readM         readManager
	dt            pbs.DataType
	pt            pbs.PacketType
	supportBinary bool
	rawReader     byteReader
	limitReader   io.LimitedReader
	base64Reader  io.Reader
}

func (pd *payloadDecoder) NextReader() (pbs.DataType, pbs.PacketType, io.ReadCloser, error) {
	if pd.rawReader == nil {
		r, supportBinary, err := pd.readM.getReader()
		if err != nil {
			return 0, 0, nil, err
		}
		br, ok := r.(byteReader)
		if !ok {
			br = bufio.NewReader(r)
		}
		if err := pd.setNextReader(br, supportBinary); err != nil {
			return 0, 0, nil, pd.sendError(err)
		}
	}
	return pd.dt, pd.pt, pd, nil
}

func (pd *payloadDecoder) Close() error {
	if _, err := io.Copy(ioutil.Discard, pd); err != nil {
		return pd.sendError(err)
	}
	err := pd.setNextReader(pd.rawReader, pd.supportBinary)
	if err != nil {
		if err != io.EOF {
			return pd.sendError(err)
		}
		pd.rawReader = nil
		pd.limitReader.R = nil
		pd.limitReader.N = 0
		pd.base64Reader = nil
		err = pd.sendError(nil)
	}
	return err
}

func (pd *payloadDecoder) Read(p []byte) (int, error) {
	if pd.base64Reader != nil {
		return pd.base64Reader.Read(p)
	}
	return pd.limitReader.Read(p)
}

func (pd *payloadDecoder) readBinary(r byteReader) (pbs.DataType, pbs.PacketType, int64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	if b > 1 {
		return 0, 0, 0, errInvalidPayload
	}
	dt := byteToDataType(b)
	l, err := readBinaryLen(r)
	if err != nil {
		return 0, 0, 0, err
	}
	b, err = r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	pt := byteToPacketType(b, dt)
	l--
	return dt, pt, l, nil
}

func (pd *payloadDecoder) readText(r byteReader) (pbs.DataType, pbs.PacketType, int64, error) {
	l, err := readTextLen(r)
	if err != nil {
		return 0, 0, 0, err
	}
	dt := pbs.DataType_string
	b, err := r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	l--
	if b == 'b' {
		dt = pbs.DataType_binary
		b, err = r.ReadByte()
		if err != nil {
			return 0, 0, 0, err
		}
		l--
	}
	pt := byteToPacketType(b, pbs.DataType_string)
	return dt, pt, l, nil
}

func (pd *payloadDecoder) sendError(err error) error {
	if e := pd.readM.addReader(err); e != nil {
		return e
	}
	return err
}

func (pd *payloadDecoder) setNextReader(r byteReader, supportBinary bool) error {
	var read func(byteReader) (pbs.DataType, pbs.PacketType, int64, error)
	if supportBinary {
		read = pd.readBinary
	} else {
		read = pd.readText
	}
	dt, pt, l, err := read(r)
	if err != nil {
		return err
	}
	pd.dt = dt
	pd.pt = pt
	pd.rawReader = r
	pd.limitReader.R = r
	pd.limitReader.N = l
	pd.supportBinary = supportBinary
	if !supportBinary && dt == pbs.DataType_binary {
		pd.base64Reader = base64.NewDecoder(base64.StdEncoding, &pd.limitReader)
	} else {
		pd.base64Reader = nil
	}
	return nil
}
