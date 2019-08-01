package socket

import (
	"encoding/json"
	"io"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

const isJSON = false

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

	if isJSON == true {
		wr := writer{
			w: w,
		}
		err := json.NewEncoder(&wr).Encode(params)
		return wr.i, err
	} else {
		// bs := []byte(params.String())
		// var err error
		bs, err := params.Marshal()
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
}

type reader struct {
	i   int64
	r   io.Reader
	buf []byte
}

func (r *reader) Read(p []byte) (n int, err error) {
	len, e := r.r.Read(p)
	r.i += int64(len)
	return len, e
}

func readConnParams(r io.Reader) (pbs.ConnParameters, error) {
	if isJSON == true {
		var params pbs.ConnParameters
		if err := json.NewDecoder(r).Decode(&params); err != nil {
			return pbs.ConnParameters{}, err
		}
		return pbs.ConnParameters{
			PingInterval: params.PingInterval,
			PingTimeout:  params.PingTimeout,
			SID:          params.SID,
			Upgrades:     params.Upgrades,
		}, nil
	} else {
		defer func() {
			if closer, ok := r.(io.Closer); ok {
				closer.Close()
			}
		}()
		rr := &reader{
			r:   r,
			buf: make([]byte, 1024*1024),
		}
		len, e := rr.Read(rr.buf)
		if e != nil {
			return pbs.ConnParameters{}, e
		}
		param := pbs.ConnParameters{}
		buf := rr.buf[:len]
		e = param.Unmarshal(buf)
		if e != nil {
			return pbs.ConnParameters{}, e
		}
		return param, nil
	}
}
