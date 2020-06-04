package packet

import (
	"io"

	"github.com/znk_fullstack/server/study/skio/ngxio/base"
)

type encoder struct {
	w FrameWriter
}

func newEncoder(w FrameWriter) *encoder {
	return &encoder{
		w: w,
	}
}

func (e *encoder) NextWriter(ft base.FrameType, pt base.PacketType) (io.WriteCloser, error) {
	w, err := e.w.NextWriter(ft)
	if err != nil {
		return nil, err
	}
	var b [1]byte
	if ft == base.FrameString {
		b[0] = pt.StringByte()
	} else {
		b[0] = pt.BinaryByte()
	}
	if _, err := w.Write(b[:]); err != nil {
		return nil, err
	}
	return w, nil
}
