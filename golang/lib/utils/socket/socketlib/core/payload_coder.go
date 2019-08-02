package core

import (
	"bytes"
	"encoding/base64"
	"io"

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
