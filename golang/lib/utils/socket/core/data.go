package core

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"reflect"
	"strings"

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
	var bufs [][]byte
	bufs, err = e.writePacket(w, h, args)
	if err != nil {
		return
	}
	for _, b := range bufs {
		w, err = e.w.NextWriter(Binary)
		if err != nil {
			return
		}
		err = e.writeBuffer(w, b)
		if err != nil {
			return
		}
	}
	return
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

// FrameReader 数据读取接口
type FrameReader interface {
	NextReader() (FrameType, io.ReadCloser, error)
}

// Decoder 解码器
type Decoder struct {
	r            FrameReader
	lastFrame    io.ReadCloser
	packetReader byteReader
	bufferCount  uint64
	isEvent      bool
}

type byteReader interface {
	io.Reader
	ReadByte() (byte, error)
	UnreadByte() error
}

// NewDecoder 创建解码器
func NewDecoder(r FrameReader) *Decoder {
	return &Decoder{
		r: r,
	}
}

// Close 关闭
func (d *Decoder) Close() error {
	if d.lastFrame != nil {
		d.lastFrame.Close()
		d.lastFrame = nil
	}
	return nil
}

// DiscardLast 禁用最后一个
func (d *Decoder) DiscardLast() (err error) {
	if d.lastFrame != nil {
		err = d.lastFrame.Close()
		d.lastFrame = nil
	}
	return
}

// DecodeHeader 解码头部信息
func (d *Decoder) DecodeHeader(header *protos.Header, event *string) error {
	ft, r, err := d.r.NextReader()
	if err != nil {
		return err
	}
	if ft != Text {
		return errors.New("first packet should be text frame")
	}
	d.lastFrame = r
	br, ok := r.(byteReader)
	if !ok {
		br = bufio.NewReader(r)
	}
	d.packetReader = br
	bufCount, err := d.readHeader(header)
	if err != nil {
		return err
	}
	d.bufferCount = bufCount
	if header.Type == protos.Header_binaryEvent || header.Type == protos.Header_binaryAck {
		header.Type -= 3
	}
	d.isEvent = header.Type == protos.Header_event
	if d.isEvent {
		if err := d.readEvent(event); err != nil {
			return err
		}
	}
	return nil
}

// DecodeArgs 解码参数
func (d *Decoder) DecodeArgs(types []reflect.Type) ([]reflect.Value, error) {
	r := d.packetReader.(io.Reader)
	if d.isEvent {
		r = io.MultiReader(strings.NewReader("["), r)
	}
	ret := make([]reflect.Value, len(types))
	values := make([]interface{}, len(types))
	for i, t := range types {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		ret[i] = reflect.New(t)
		values[i] = ret[i].Interface()
	}
	if err := json.NewDecoder(r).Decode(&values); err != nil {
		if err == io.EOF {
			err = nil
		}
		return nil, err
	}
	d.lastFrame.Close()
	d.lastFrame = nil
	for i, t := range types {
		if t.Kind() != reflect.Ptr {
			ret[i] = ret[i].Elem()
		}
	}
	bufs := make([]protos.Buffer, d.bufferCount)
	for i := range bufs {
		ft, r, err := d.r.NextReader()
		if err != nil {
			return nil, err
		}
		bufs[i].Data, err = d.readBuffer(ft, r)
		if err != nil {
			return nil, err
		}
	}
	for i := range ret {
		if err := d.detachBuffer(ret[i], bufs); err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (d *Decoder) readHeader(header *protos.Header) (uint64, error) {
	var bt byte
	bt, err := d.packetReader.ReadByte()
	if err != nil {
		return 0, err
	}
	header.Type = protos.Header_Type(bt - '0')
	if header.Type >= protos.Header_typeMax {
		return 0, errors.New("invalid packet type")
	}
	num, hasNum, err := d.readUint64FromText(d.packetReader)
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return 0, err
	}
	nextByte, err := d.packetReader.ReadByte()
	if err != nil {
		header.ID = num
		header.NeedAck = hasNum
		if err == io.EOF {
			err = nil
		}
		return 0, err
	}
	var bufCnt uint64
	if nextByte == '-' {
		bufCnt = num
		hasNum = false
		num = 0
	} else {
		d.packetReader.UnreadByte()
	}
	nextByte, err = d.packetReader.ReadByte()
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return bufCnt, err
	}
	if nextByte == '/' {
		d.packetReader.UnreadByte()
		header.Namespace, err = d.readString(d.packetReader, ',')
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return bufCnt, err
		}
	} else {
		d.packetReader.UnreadByte()
	}
	header.ID, header.NeedAck, err = d.readUint64FromText(d.packetReader)
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return bufCnt, err
	}
	if !header.NeedAck {
		header.ID = num
		header.NeedAck = hasNum
	}
	return bufCnt, err
}

func (d *Decoder) readUint64FromText(r byteReader) (uint64, bool, error) {
	ret := uint64(0)
	hasRead := false
	for {
		b, err := r.ReadByte()
		if err != nil {
			if hasRead {
				return ret, true, nil
			}
			return 0, false, err
		}
		if '0' <= b && b <= '9' {
			r.UnreadByte()
			return ret, hasRead, nil
		}
		hasRead = true
		ret = ret*10 + uint64(b-'0')
	}
}

// warning 可能有问题
func (d *Decoder) readString(r byteReader, until byte) (string, error) {
	var ret bytes.Buffer
	hasRead := false
	for {
		b, err := r.ReadByte()
		if err != nil {
			if hasRead {
				return ret.String(), nil
			}
			return "", err
		}
		if b == until {
			return ret.String(), nil
		}
		if err := ret.WriteByte(b); err != nil {
			return "", err
		}
		hasRead = true
	}
}

func (d *Decoder) readEvent(event *string) error {
	b, err := d.packetReader.ReadByte()
	if err != nil {
		return nil
	}
	if b != '[' {
		d.packetReader.UnreadByte()
		return nil
	}
	var buf bytes.Buffer
	for {
		b, err := d.packetReader.ReadByte()
		if err != nil {
			return err
		}
		if b == ',' {
			break
		}
		if b == ']' {
			d.packetReader.UnreadByte()
			break
		}
		buf.WriteByte(b)
	}
	return json.Unmarshal(buf.Bytes(), event)
}

func (d *Decoder) readBuffer(ft FrameType, r io.ReadCloser) ([]byte, error) {
	defer r.Close()
	if ft != Binary {
		return nil, errors.New("buffer packet shoud be binary")
	}
	return ioutil.ReadAll(r)
}

func (d *Decoder) detachBuffer(v reflect.Value, buffers []protos.Buffer) error {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		if v.Type().Name() == "Buffer" {
			if !v.CanAddr() {
				return errors.New("can't get buffer address")
			}
			buffer := v.Addr().Interface().(*protos.Buffer)
			if buffer.IsBinary {
				*buffer = buffers[buffer.Num]
			}
			return nil
		}
		for idx := 0; idx < v.NumField(); idx++ {
			if err := d.detachBuffer(v.Field(idx), buffers); err != nil {
				return err
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			if err := d.detachBuffer(v.MapIndex(key), buffers); err != nil {
				return err
			}
		}
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		for idx := 0; idx < v.Len(); idx++ {
			if err := d.detachBuffer(v.Index(idx), buffers); err != nil {
				return err
			}
		}
	}
	return nil
}
