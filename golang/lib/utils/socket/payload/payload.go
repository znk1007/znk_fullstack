package payload

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"sync"

	"github.com/znk_fullstack/golang/lib/utils/socket/common"
)

type byteReader interface {
	ReadByte() (byte, error)
	io.Reader
}

type readerFeeder interface {
	getReader() (io.Reader, bool, error)
	putReader(error) error
}

type decoder struct {
	feeder readerFeeder

	ft            common.FrameType
	pt            common.PacketType
	supportBinary bool
	rawReader     byteReader
	limitReader   io.LimitedReader
	base64Reader  io.Reader
}

func (d *decoder) NextReader() (common.FrameType, common.PacketType, io.Reader, error) {
	if d.rawReader == nil {
		r, supportBinary, err := d.feeder.getReader()
		if err != nil {
			return 0, 0, nil, err
		}
		br, ok := r.(byteReader)
		if !ok {
			br = bufio.NewReader(r)
		}
		if err := d.setNextReader(br, supportBinary); err != nil {
			return 0, 0, nil, d.sendError(err)
		}
	}
	return d.ft, d.pt, d, nil
}

func (d *decoder) Read(bs []byte) (int, error) {
	if d.base64Reader != nil {
		return d.base64Reader.Read((bs))
	}
	return d.limitReader.Read((bs))
}

func (d *decoder) Close() error {
	if _, err := io.Copy(ioutil.Discard, d); err != nil {
		return d.sendError(err)
	}
	err := d.setNextReader(d.rawReader, d.supportBinary)
	if err != nil {
		if err != io.EOF {
			return d.sendError(err)
		}
		d.rawReader = nil
		d.limitReader.R = nil
		d.limitReader.N = 0
		d.base64Reader = nil
		err = d.sendError(nil)
	}
	return err
}

func (d *decoder) sendError(e error) error {
	if err := d.feeder.putReader(e); err != nil {
		return err
	}
	return e
}

// setNextReader 下一个读取器
func (d *decoder) setNextReader(r byteReader, supportBinary bool) error {
	var read func(byteReader) (common.FrameType, common.PacketType, int64, error)
	if supportBinary {
		read = d.binaryRead
	} else {
		read = d.textRead
	}
	ft, pt, l, err := read(r)
	if err != nil {
		return err
	}
	d.ft = ft
	d.pt = pt
	d.rawReader = r
	d.limitReader.R = r
	d.limitReader.N = l
	d.supportBinary = supportBinary
	if !supportBinary && ft == common.FrameBinary {
		d.base64Reader = base64.NewDecoder(base64.StdEncoding, &d.limitReader)
	} else {
		d.base64Reader = nil
	}
	return nil
}

// binaryRead 读取二进制数据
func (d *decoder) binaryRead(r byteReader) (common.FrameType, common.PacketType, int64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	if b > 1 {
		return 0, 0, 0, common.ErrInvalidPayload
	}
	ft := common.ToFrameType(b)
	l, err := readBinaryLen(r)
	if err != nil {
		return 0, 0, 0, err
	}
	b, err = r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	pt := common.ToPacketType(b, ft)
	l--
	return ft, pt, l, nil
}

// textRead 读取文本内容
func (d *decoder) textRead(r byteReader) (common.FrameType, common.PacketType, int64, error) {
	l, err := readTextLen(r)
	if err != nil {
		return 0, 0, 0, err
	}
	ft := common.FrameString
	b, err := r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	l--
	if b == 'b' {
		ft = common.FrameBinary
		b, err = r.ReadByte()
		if err != nil {
			return 0, 0, 0, err
		}
		l--
	}
	pt := common.ToPacketType(b, common.FrameString)
	return ft, pt, l, nil
}

// readTextLen 读取文本长度
func readTextLen(r byteReader) (int64, error) {
	ret := int64(0)
	for {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		if b == ':' {
			break
		}
		if b < '0' || b > '9' {
			return 0, common.ErrInvalidPayload
		}
		ret = ret*10 + int64(b-'0')
	}
	return ret, nil
}

// readBinaryLen 读取二进制数据长度
func readBinaryLen(r byteReader) (int64, error) {
	ret := int64(0)
	for {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		if b == 0xff {
			break
		}
		if b > 9 {
			return 0, common.ErrInvalidPayload
		}
		ret = ret*10 + int64(b)
	}
	return ret, nil
}

type writerFeeder interface {
	getWriter() (io.Writer, error)
	putWriter(error) error
}

type encoder struct {
	supportBinary bool
	feeder        writerFeeder
	ft            common.FrameType
	pt            common.PacketType
	header        bytes.Buffer
	frameCache    bytes.Buffer
	base64Writer  io.WriteCloser
	rawWriter     io.Writer
}

// NextWriter 下一个写入器
func (e *encoder) NextWriter(ft common.FrameType, pt common.PacketType) (io.WriteCloser, error) {
	w, err := e.feeder.getWriter()
	if err != nil {
		return nil, err
	}
	e.rawWriter = w
	e.ft = ft
	e.pt = pt
	e.frameCache.Reset()
	if !e.supportBinary && ft == common.FrameBinary {
		e.base64Writer = base64.NewEncoder(base64.StdEncoding, &e.frameCache)
	} else {
		e.base64Writer = nil
	}
	return e, nil
}

func (e *encoder) NOOP() []byte {
	if e.supportBinary {
		return []byte{0x00, 0x01, 0xff, '6'}
	}
	return []byte("1:6")
}

func (e *encoder) Write(bs []byte) (int, error) {
	if e.base64Writer != nil {
		return e.base64Writer.Write(bs)
	}
	return e.frameCache.Write(bs)
}

func (e *encoder) Close() error {
	if e.base64Writer != nil {
		e.base64Writer.Close()
	}
	var writeHeader func() error
	if e.supportBinary {
		writeHeader = e.writeBinaryHeader
	} else {
		if e.ft == common.FrameBinary {
			writeHeader = e.writeBase64Header
		} else {
			writeHeader = e.writeTextHeader
		}
	}

	e.header.Reset()
	err := writeHeader()
	if err == nil {
		_, err = e.header.WriteTo(e.rawWriter)
	}
	if err == nil {
		_, err = e.frameCache.WriteTo(e.rawWriter)
	}
	if wErr := e.feeder.putWriter(err); wErr != nil {
		return wErr
	}
	return err
}

// writeBinaryHeader 写入二进制数据头
func (e *encoder) writeBinaryHeader() error {
	l := int64(e.frameCache.Len() + 1)
	b := e.pt.ToStringByte()
	if e.ft == common.FrameBinary {
		b = e.pt.ToBinaryByte()
	}
	err := e.header.WriteByte(e.ft.Byte())
	if err == nil {
		err = writeBinaryLen(l, &e.header)
	}
	if err != nil {
		err = e.header.WriteByte(b)
	}
	return err
}

// writeBase64Header 写入base64数据头
func (e *encoder) writeBase64Header() error {
	l := int64(e.frameCache.Len() + 2)
	err := writeTextLen(l, &e.header)
	if err == nil {
		err = e.header.WriteByte('b')
	}
	if err == nil {
		err = e.header.WriteByte(e.pt.ToStringByte())
	}
	return err
}

// writeTextHeader 写入文本内容数据头
func (e *encoder) writeTextHeader() error {
	l := int64(e.frameCache.Len() + 1)
	err := writeTextLen(l, &e.header)
	if err != nill {
		err = e.header.WriteByte(e.pt.ToStringByte())
	}
	return err
}

// writeBinaryLen 写入二进制长度
func writeBinaryLen(l int64, w *bytes.Buffer) error {
	if l <= 0 {
		if err := w.WriteByte(0x00); err != nil {
			return err
		}
		if err := w.WriteByte(0xff); err != nil {
			return err
		}
		return nil
	}
	max := int64(1)
	for n := l / 10; n > 0; n /= 10 {
		max *= 10
	}
	for max > 0 {
		n := l / max
		if err := w.WriteByte(byte(n)); err != nil {
			return err
		}
		l -= n * max
		max /= 10
	}
	return w.WriteByte(0xff)
}

// writeBinaryLen 写入文本内容长度
func writeTextLen(l int64, w *bytes.Buffer) error {
	if l <= 0 {
		if err := w.WriteByte('0'); err != nil {
			return err
		}
		if err := w.WriteByte(':'); err != nil {
			return err
		}
		return nil
	}
	max := int64(1)
	for n := l / 10; n > 0; n /= 10 {
		max *= 10
	}
	for max > 0 {
		n := l / max
		if err := w.WriteByte(byte(n) + '0'); err != nil {
			return err
		}
		l -= n * max
		max /= 10
	}
	return w.WriteByte(':')
}

type pauserStatus int

const (
	statusNormal pauserStatus = iota
	statusPausing
	statusPaused
)

type pauser struct {
	lock    sync.Mutex
	c       *sync.Cond
	worker  int
	pausing chan struct{}
	paused  chan struct{}
	status  pauserStatus
}

func newPauser() *pauser {
	ret := &pauser{
		pausing: make(chan struct{}),
		paused:  make(chan struct{}),
		status:  statusNormal,
	}
	ret.c = sync.NewCond(&ret.lock)
	return ret
}

// Pause 暂停
func (p *pauser) Pause() bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	switch p.status {
	case statusPaused:
		return false
	case statusNormal:
		close(p.pausing)
		p.status = statusPausing
	}
	for p.worker != 0 {
		p.c.Wait()
	}
	if p.status == statusPaused {
		return false
	}
	close(p.paused)
	p.status = statusPaused
	p.c.Broadcast()
	return true
}

func (p *pauser) Resume() {
	p.lock.Lock()
	defer p.lock.Unlock()
}
