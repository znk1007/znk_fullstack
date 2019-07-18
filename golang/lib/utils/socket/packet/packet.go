package packet

import (
	"io"

	"github.com/znk1007/fullstack/lib/utils/socket/common"
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

}
