package socket

import (
	"fmt"
	"io"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

type writer struct {
	i int64
	w io.Writer
}

func (w *writer) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	w.i += int64(n)
	return n, err
}

// writeTo 写入数据
func writeTo(params pbs.ConnParameters, w io.Writer) (int64, error) {

	bs := []byte(params.String())
	var err error
	// bs, err := params.Marshal()
	if err != nil {
		return 0, err
	}
	wr := &writer{
		w: w,
	}
	_, err = wr.Write(bs)
	if err != nil {
		return 0, err
	}
	return wr.i, nil
}

type reader struct {
	i   int64
	r   io.Reader
	buf []byte
}

func (r *reader) Read(p []byte) (n int, err error) {
	len, e := r.r.Read(p)
	r.i += int64(len)
	return n, e
}

func readConnParams(r io.Reader) (*pbs.ConnParameters, error) {
	defer func() {
		if closer, ok := r.(io.Closer); ok {
			closer.Close()
		}
	}()
	rr := &reader{
		r:   r,
		buf: make([]byte, 1024*1024),
	}
	_, e := rr.Read(rr.buf)
	if e != nil {
		return nil, e
	}
	param := &pbs.ConnParameters{}
	e = param.Unmarshal(rr.buf)
	fmt.Println("unmarshal param: ", param)
	if e != nil {
		return nil, e
	}
	return param, nil
}
