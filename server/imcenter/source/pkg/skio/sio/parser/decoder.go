package parser

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/znk_fullstack/server/study/skio/ngxio"
)

//FrameReader read data frame interface
type FrameReader interface {
	NextReader() (ngxio.FrameType, io.ReadCloser, error)
}

//Decoder decode data from transport
type Decoder struct {
	r FrameReader

	lastFrame    io.ReadCloser
	packetReader byteReader
	bufferCount  uint64
	isEvent      bool
}

//NewDecoder new the data decode struct
func NewDecoder(r FrameReader) *Decoder {
	return &Decoder{
		r: r,
	}
}

//Close close the frame from decoder
func (d *Decoder) Close() error {
	if d.lastFrame != nil {
		d.lastFrame.Close()
		d.lastFrame = nil
	}
	return nil
}

type byteReader interface {
	io.Reader
	ReadByte() (byte, error)
	UnreadByte() error
}

//DiscardLast discard the last data frame
func (d *Decoder) DiscardLast() (err error) {
	if d.lastFrame != nil {
		err = d.lastFrame.Close()
		d.lastFrame = nil
	}
	return
}

//DecodeHeader decode the data header frame
func (d *Decoder) DecodeHeader(header *Header, event *string) error {
	ft, r, err := d.r.NextReader()
	if err != nil {
		return err
	}
	if ft != ngxio.TEXT {
		return errors.New("first packet should be TEXT frame")
	}
	d.lastFrame = r
	br, ok := r.(byteReader)
	if !ok {
		br = bufio.NewReader(r)
	}
	d.packetReader = br

	bufCnt, err := d.readHeader(header)
	if err != nil {
		return err
	}
	d.bufferCount = bufCnt
	if header.Type == binaryEvent || header.Type == binaryAck {
		header.Type -= 3
	}
	d.isEvent = header.Type == Event
	if d.isEvent {
		if err := d.readEvent(event); err != nil {
			return err
		}
	}
	return nil
}

//DecodeArgs decode the arg parameters
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
		d.DiscardLast()
		return nil, err
	}
	//we can't use defer or call DiscardLast before decoding, because
	//there are buffered readers involved and if we invoke .Close() json
	//will encounter upexpected EOF.
	d.DiscardLast()

	for i, t := range types {
		if t.Kind() != reflect.Ptr {
			ret[i] = ret[i].Elem()
		}
	}

	bufs := make([]Buffer, d.bufferCount)
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
		if !('0' <= b && b <= '9') {
			r.UnreadByte()
			return ret, hasRead, nil
		}
		hasRead = true
		ret = ret*10 + uint64(b-'0')
	}
}

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

func (d *Decoder) readHeader(header *Header) (uint64, error) {
	//read type
	var bt byte
	bt, err := d.packetReader.ReadByte()
	if err != nil {
		return 0, err
	}
	header.Type = Type(bt - '0')
	if header.Type >= typeMax {
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
	//check if buffer count
	var bufferCount uint64
	if nextByte == '-' {
		bufferCount = num
		hasNum = false
		num = 0
	} else {
		d.packetReader.UnreadByte()
	}

	//check namespace
	nextByte, err = d.packetReader.ReadByte()
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return bufferCount, err
	}
	if nextByte == '/' {
		d.packetReader.UnreadByte()
		header.Namespace, err = d.readString(d.packetReader, ',')
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return bufferCount, err
		}
	} else {
		d.packetReader.UnreadByte()
	}

	//read id
	header.ID, header.NeedAck, err = d.readUint64FromText(d.packetReader)
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return bufferCount, err
	}
	if !header.NeedAck {
		//313["data"], id has beed read at beginning, need add back.
		header.ID = num
		header.NeedAck = hasNum
	}
	return bufferCount, err
}

func (d *Decoder) readEvent(event *string) error {
	b, err := d.packetReader.ReadByte()
	if err != nil {
		return err
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

func (d *Decoder) readBuffer(ft ngxio.FrameType, r io.ReadCloser) ([]byte, error) {
	defer r.Close()
	if ft != ngxio.BINARY {
		return nil, errors.New("buffer packet should be BINARY")
	}
	return ioutil.ReadAll(r)
}

func (d *Decoder) detachBuffer(v reflect.Value, bufs []Buffer) error {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		if v.Type().Name() == "Buffer" {
			if !v.CanAddr() {
				return errors.New("can't get Buffer address")
			}
			buffer := v.Addr().Interface().(*Buffer)
			if buffer.isBinary {
				*buffer = bufs[buffer.num]
			}
			return nil
		}
		for i := 0; i < v.NumField(); i++ {
			if err := d.detachBuffer(v.Field(i), bufs); err != nil {
				return err
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			if err := d.detachBuffer(v.MapIndex(key), bufs); err != nil {
				return err
			}
		}
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if err := d.detachBuffer(v.Index(i), bufs); err != nil {
				return err
			}
		}
	}
	return nil
}
