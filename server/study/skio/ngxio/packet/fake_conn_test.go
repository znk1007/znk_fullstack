package packet

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/znk_fullstack/server/study/skio/ngxio/base"
)

type fakeFrame struct {
	w    *fakeConnWriter
	ft   base.FrameType
	data *bytes.Buffer
}

func newFakeFrame(w *fakeConnWriter, ft base.FrameType) *fakeFrame {
	return &fakeFrame{
		w:    w,
		ft:   ft,
		data: bytes.NewBuffer(nil),
	}
}

func (w *fakeFrame) Write(p []byte) (int, error) {
	return w.data.Write(p)
}

func (w *fakeFrame) Read(p []byte) (int, error) {
	return w.data.Read(p)
}

func (w *fakeFrame) Close() error {
	if w.w == nil {
		return nil
	}
	w.w.frames = append(w.w.frames, Frame{
		ft:   w.ft,
		data: w.data.Bytes(),
	})
	return nil
}

type fakeConnReader struct {
	frames []Frame
}

func newFakeConnReader(frames []Frame) *fakeConnReader {
	return &fakeConnReader{
		frames: frames,
	}
}

func (fcr *fakeConnReader) NextReader() (base.FrameType, io.ReadCloser, error) {
	if len(fcr.frames) == 0 {
		return base.FrameString, nil, io.EOF
	}
	f := fcr.frames[0]
	fcr.frames = fcr.frames[1:]
	return f.ft, ioutil.NopCloser(bytes.NewReader(f.data)), nil
}

type fakeConnWriter struct {
	frames []Frame
}

func newFakeConnWriter() *fakeConnWriter {
	return &fakeConnWriter{}
}

func (w *fakeConnWriter) NextWriter(ft base.FrameType) (io.WriteCloser, error) {
	return newFakeFrame(w, ft), nil
}

type fakeOneFrameConst struct {
	b byte
}

func (c *fakeOneFrameConst) Read(p []byte) (int, error) {
	p[0] = c.b
	return 1, nil
}

type fakeConstReader struct {
	ft base.FrameType
	r  *fakeOneFrameConst
}

func newFakeConstReader() *fakeConstReader {
	return &fakeConstReader{
		ft: base.FrameString,
		r: &fakeOneFrameConst{
			b: base.MESSAGE.StringByte(),
		},
	}
}

func (fcr *fakeConstReader) NextReader() (base.FrameType, io.ReadCloser, error) {
	ft := fcr.ft
	switch ft {
	case base.FrameBinary:
		fcr.ft = base.FrameString
		fcr.r.b = base.MESSAGE.StringByte()
	case base.FrameString:
		fcr.ft = base.FrameBinary
		fcr.r.b = base.MESSAGE.BinaryByte()
	}
	return ft, ioutil.NopCloser(fcr.r), nil
}

type fakeOneFrameDiscarder struct{}

func (d fakeOneFrameDiscarder) Write(p []byte) (int, error) {
	return len(p), nil
}

func (d fakeOneFrameDiscarder) Close() error {
	return nil
}

type fakeDiscardWriter struct{}

func (w *fakeDiscardWriter) NextWriter(ft base.FrameType) (io.WriteCloser, error) {
	return fakeOneFrameDiscarder{}, nil
}
