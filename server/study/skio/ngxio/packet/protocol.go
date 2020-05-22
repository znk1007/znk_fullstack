package packet

import (
	"io"

	"github.com/znk_fullstack/server/study/skio/ngxio/base"
)

//FrameReader is the reader which supports framing.
type FrameReader interface {
	NextReader() (base.FrameType, io.ReadCloser, error)
}

//FrameWriter is the writer which supports framing.
type FrameWriter interface {
	NextWriter(ft base.FrameType) (io.WriteCloser, error)
}

//NewEncoder creates a packet encoder which writes to w.
func NewEncoder(w FrameWriter) base.FrameWriter {
	return newEncoder(w)
}

//NewDecoder creates a packet decoder which reads from r.
func NewDecoder(r FrameReader) base.FrameReader {
	return newDecoder(r)
}
