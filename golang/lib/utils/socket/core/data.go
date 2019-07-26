package core

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"reflect"

	_ "github.com/gogo/protobuf/io"
	protos "github.com/znk_fullstack/golang/lib/utils/socket/protos/generated"
)

type byteWriter interface {
	io.Writer
	WriteByte(byte) error
}

type flusher interface {
	Flush() error
}

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

func (e *Encoder) writeUint64(w byteWriter, i uint64) error {
	base := uint64(1)
	for i/base > 10 {
		base *= 10
	}
	for base > 0 {
		p := i / base
		if err := w.WriteByte(byte(p) + '0'); err != nil {
			return err
		}
		i -= p * base
		base /= 10
	}
	return nil
}

func (e *Encoder) writePacket(w io.WriteCloser, h protos.Header, args []interface{}) ([][]byte, error) {
	defer w.Close()
	bw, ok := w.(byteWriter)
	if !ok {
		bw = bufio.NewWriter(w)
	}
	max := uint64(0)
	bufs, err := e.attachBuffer(reflect.ValueOf(args), &max)
	if err != nil {
		return nil, err
	}
	if len(bufs) > 0 && (h.Type == protos.Header_event || h.Type == protos.Header_ack) {
		h.Type += 3
	}
	if err := bw.WriteByte(byte(h.Type + '0')); err != nil {
		return nil, err
	}
	if h.Type == protos.Header_binaryAck || h.Type == protos.Header_binaryEvent {
		if err := e.writeUint64(bw, max); err != nil {
			return nil, err
		}
		if err := bw.WriteByte('-'); err != nil {
			return nil, err
		}
	}
	if h.Namespace != "" {
		if _, err := bw.Write([]byte(h.Namespace)); err != nil {
			return nil, err
		}
		if h.ID != 0 || args != nil {
			if err := bw.WriteByte(','); err != nil {
				return nil, err
			}
		}
	}

	if h.NeedAck {
		if err := e.writeUint64(bw, h.ID); err != nil {
			return nil, err
		}
	}
	if args != nil {
		// fullWriter := gogoio.NewFullWriter(bw).WriteMsg(args)
		if err := json.NewEncoder(bw).Encode(args); err != nil {
			return nil, err
		}
	}
	if f, ok := bw.(flusher); ok {
		if err := f.Flush(); err != nil {
			return nil, err
		}
	}
	return bufs, nil
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
