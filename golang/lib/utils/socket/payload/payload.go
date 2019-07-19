package payload

import (
	"bufio"
	"encoding/base64"
	"io"
	"io/ioutil"

	"github.com/znk_fullstack/golang/lib/utils/socket/common"
)

type byteReader interface {
	ReadByte() (byte, error)
	io.Reader
}

type readerFeeder interface {
	getReader() (io.Reader, bool, error)
	putReader(error) error
}

type decoder struct {
	feeder readerFeeder

	ft            common.FrameType
	pt            common.PacketType
	supportBinary bool
	rawReader     byteReader
	limitReader   io.LimitedReader
	base64Reader  io.Reader
}

func (d *decoder) NextReader() (common.FrameType, common.PacketType, io.Reader, error) {
	if d.rawReader == nil {
		r, supportBinary, err := d.feeder.getReader()
		if err != nil {
			return 0, 0, nil, err
		}
		br, ok := r.(byteReader)
		if !ok {
			br = bufio.NewReader(r)
		}
		if err := d.setNextReader(br, supportBinary); err != nil {
			return 0, 0, nil, d.sendError(err)
		}
	}
	return d.ft, d.pt, d, nil
}

func (d *decoder) Read(bs []byte) (int, error) {
	if d.base64Reader != nil {
		return d.base64Reader.Read((bs))
	}
	return d.limitReader.Read((bs))
}

func (d *decoder) Close() error {
	if _, err := io.Copy(ioutil.Discard, d); err != nil {
		return d.sendError(err)
	}
	err := d.setNextReader(d.rawReader, d.supportBinary)
	if err != nil {
		if err != io.EOF {
			return d.sendError(err)
		}
		d.rawReader = nil
		d.limitReader.R = nil
		d.limitReader.N = 0
		d.base64Reader = nil
		err = d.sendError(nil)
	}
	return err
}

func (d *decoder) sendError(e error) error {
	if err := d.feeder.putReader(e); err != nil {
		return err
	}
	return e
}

// setNextReader 下一个读取器
func (d *decoder) setNextReader(r byteReader, supportBinary bool) error {
	var read func(byteReader) (common.FrameType, common.PacketType, int64, error)
	if supportBinary {
		read = d.binaryRead
	} else {
		read = d.textRead
	}
	ft, pt, l, err := read(r)
	if err != nil {
		return err
	}
	d.ft = ft
	d.pt = pt
	d.rawReader = r
	d.limitReader.R = r
	d.limitReader.N = l
	d.supportBinary = supportBinary
	if !supportBinary && ft == common.FrameBinary {
		d.base64Reader = base64.NewDecoder(base64.StdEncoding, &d.limitReader)
	} else {
		d.base64Reader = nil
	}
	return nil
}

// binaryRead 读取二进制数据
func (d *decoder) binaryRead(r byteReader) (common.FrameType, common.PacketType, int64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	if b > 1 {
		return 0, 0, 0, common.ErrInvalidPayload
	}
	ft := common.ToFrameType(b)
	l, err := readBinaryLen(r)
	if err != nil {
		return 0, 0, 0, err
	}
	b, err = r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	pt := common.ToPacketType(b, ft)
	l--
	return ft, pt, l, nil
}

// textRead 读取文本内容
func (d *decoder) textRead(r byteReader) (common.FrameType, common.PacketType, int64, error) {
	l, err := readTextLen(r)
	if err != nil {
		return 0, 0, 0, err
	}
	ft := common.FrameString
	b, err := r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	l--
	if b == 'b' {
		ft = common.FrameBinary
		b, err = r.ReadByte()
		if err != nil {
			return 0, 0, 0, err
		}
		l--
	}
	pt := common.ToPacketType(b, common.FrameString)
	return ft, pt, l, nil
}

// readTextLen 读取文本长度
func readTextLen(r byteReader) (int64, error) {
	ret := int64(0)
	for {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		if b == ':' {
			break
		}
		if b < '0' || b > '9' {
			return 0, common.ErrInvalidPayload
		}
		ret = ret*10 + int64(b-'0')
	}
	return ret, nil
}

// readBinaryLen 读取二进制数据长度
func readBinaryLen(r byteReader) (int64, error) {
	ret := int64(0)
	for {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		if b == 0xff {
			break
		}
		if b > 9 {
			return 0, common.ErrInvalidPayload
		}
		ret = ret*10 + int64(b)
	}
	return ret, nil
}

type writerFeeder interface {
	getWriter() (io.Writer, error)
	putWriter(error) error
}

type encoder struct {
	supportBinary bool
	feeder        writerFeeder
	ft            common.FrameType
}
