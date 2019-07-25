package core

import (
	"errors"
	"io"
	"reflect"

	protos "github.com/znk_fullstack/golang/lib/utils/socket/protos/generated"
)

// FrameWriter 数据写入器
type FrameWriter interface {
	NextWriter(ft FrameType) (io.WriteCloser, error)
}

// Encoder 编码
type Encoder struct {
	w FrameWriter
}

// NewEncoder 新建编码器
func NewEncoder(w FrameWriter) *Encoder {
	return &Encoder{
		w: w,
	}
}

// Encode 编码
func (e *Encoder) Encode(h protos.Header, args []interface{}) (err error) {
	var w io.WriteCloser
	w, err = e.w.NextWriter(Text)
	if err != nil {
		return
	}

}

// 写入二进制数据
func (e *Encoder) writeBuffer(w io.WriteCloser, buffer []byte) error {
	defer w.Close()
	_, err := w.Write(buffer)
	return err
}

func (e *Encoder) attachBuffer(v reflect.Value, idx *uint64) ([][]byte, error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	var ret [][]byte
	switch v.Kind() {
	case reflect.Struct:
		if v.Type().Name() == "Buffer" {
			if !v.CanAddr() {
				return nil, errors.New("cannot get Buffer address")
			}
			buffer := v.Addr().Interface().(*protos.Buffer)
			buffer.Num = *idx
			buffer.IsBinary = true
			ret = append(ret, buffer.Data)
			*idx++
		} else {
			for i := 0; i < v.NumField(); i++ {
				b, err := e.attachBuffer(v.Field(i), idx)
				if err != nil {
					return nil, err
				}
				ret = append(ret, b...)
			}
		}
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			b, err := e.attachBuffer(v.Index(i), idx)
			if err != nil {
				return nil, err
			}
			ret = append(ret, b...)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			b, err := e.attachBuffer(v.MapIndex(key), idx)
			if err != nil {
				return nil, err
			}
			ret = append(ret, b...)
		}
	}
	return ret, nil
}
