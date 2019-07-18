package packet

import (
	"io"

	"github.com/znk_fullstack/golang/lib/utils/socket/common"
)

// FrameReader 数据读取器
type FrameReader interface {
	NextReader() (common.FrameType, io.ReadCloser, error)
}

// FrameWriter 数据写入器
type FrameWriter interface {
	NextWriter(ft common.FrameType) (io.WriteCloser, error)
}

// NewEncoder 创建编码器
func NewEncoder(fw FrameWriter) common.FrameWriter {
	return newEncoder(fw)
}

// NewDecoder 创建解码器
func NewDecoder(fr FrameReader) common.FrameReader {
	return newDecoder(fr)
}

// encoder 编码器
type encoder struct {
	w FrameWriter
}

func newEncoder(w FrameWriter) *encoder {
	return &encoder{
		w: w,
	}
}

// NextWriter 下一个写入器
func (e *encoder) NextWriter(ft common.FrameType, pt common.PacketType) (io.WriteCloser, error) {
	w, err := e.w.NextWriter(ft)
	if err != nil {
		return nil, err
	}
	var b [1]byte
	if ft == common.FrameString {
		b[0] = pt.ToStringByte()
	} else {
		b[0] = pt.ToBinaryByte()
	}
	if _, err := w.Write(b[:]); err != nil {
		w.Close()
		return nil, err
	}
	return w, nil
}

type decoder struct {
	r FrameReader
}

func newDecoder(r FrameReader) *decoder {
	return &decoder{
		r: r,
	}
}

func (d *decoder) NextReader() (common.FrameType, common.PacketType, io.ReadCloser, error) {
	ft, r, err := d.r.NextReader()
	if err != nil {
		return 0, 0, nil, err
	}
	var b [1]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return 0, 0, nil, err
	}
	return ft, common.ToPacketType(b[0], ft), r, nil
}
