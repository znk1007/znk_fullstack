package payload

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/znk_fullstack/golang/lib/utils/socket/primary"
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

	ft            primary.FrameType
	pt            primary.PacketType
	supportBinary bool
	rawReader     byteReader
	limitReader   io.LimitedReader
	base64Reader  io.Reader
}

func (d *decoder) NextReader() (primary.FrameType, primary.PacketType, io.ReadCloser, error) {
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
	var read func(byteReader) (primary.FrameType, primary.PacketType, int64, error)
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
	if !supportBinary && ft == primary.FrameBinary {
		d.base64Reader = base64.NewDecoder(base64.StdEncoding, &d.limitReader)
	} else {
		d.base64Reader = nil
	}
	return nil
}

// binaryRead 读取二进制数据
func (d *decoder) binaryRead(r byteReader) (primary.FrameType, primary.PacketType, int64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	if b > 1 {
		return 0, 0, 0, primary.ErrInvalidPayload
	}
	ft := primary.ToFrameType(b)
	l, err := readBinaryLen(r)
	if err != nil {
		return 0, 0, 0, err
	}
	b, err = r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	pt := primary.ToPacketType(b, ft)
	l--
	return ft, pt, l, nil
}

// textRead 读取文本内容
func (d *decoder) textRead(r byteReader) (primary.FrameType, primary.PacketType, int64, error) {
	l, err := readTextLen(r)
	if err != nil {
		return 0, 0, 0, err
	}
	ft := primary.FrameString
	b, err := r.ReadByte()
	if err != nil {
		return 0, 0, 0, err
	}
	l--
	if b == 'b' {
		ft = primary.FrameBinary
		b, err = r.ReadByte()
		if err != nil {
			return 0, 0, 0, err
		}
		l--
	}
	pt := primary.ToPacketType(b, primary.FrameString)
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
			return 0, primary.ErrInvalidPayload
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
			return 0, primary.ErrInvalidPayload
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
	ft            primary.FrameType
	pt            primary.PacketType
	header        bytes.Buffer
	frameCache    bytes.Buffer
	base64Writer  io.WriteCloser
	rawWriter     io.Writer
}

// NextWriter 下一个写入器
func (e *encoder) NextWriter(ft primary.FrameType, pt primary.PacketType) (io.WriteCloser, error) {
	w, err := e.feeder.getWriter()
	if err != nil {
		return nil, err
	}
	e.rawWriter = w
	e.ft = ft
	e.pt = pt
	e.frameCache.Reset()
	if !e.supportBinary && ft == primary.FrameBinary {
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
		if e.ft == primary.FrameBinary {
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
	if e.ft == primary.FrameBinary {
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
	if err != nil {
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

// Resume 复位
func (p *pauser) Resume() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.paused = make(chan struct{})
	p.pausing = make(chan struct{})
}

// Working 正在工作
func (p *pauser) Working() bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.status == statusPaused {
		return false
	}
	p.worker++
	return true
}

// Done 完成
func (p *pauser) Done() {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.status == statusPaused || p.worker == 0 {
		return
	}
	p.worker--
	p.c.Broadcast()
}

// PausingTrigger 即将暂停操作的管道
func (p *pauser) PausingTrigger() <-chan struct{} {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.pausing
}

// PausedTrigger 已暂停操作的管道
func (p *pauser) PausedTrigger() <-chan struct{} {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.paused
}

type readArg struct {
	r             io.Reader
	supportBinary bool
}

// Payload 有效载荷
type Payload struct {
	close     chan struct{}
	closeOnce sync.Once
	err       atomic.Value

	pauser *pauser

	readerChan   chan readArg
	feeding      int64
	readError    chan error
	readDeadline atomic.Value
	decoder      decoder

	writerChan    chan io.Writer
	flushing      int64
	writeError    chan error
	writeDeadline atomic.Value
	encoder       encoder
}

// New 新建有效载荷
func New(supportBinary bool) *Payload {
	ret := &Payload{
		close:      make(chan struct{}),
		pauser:     newPauser(),
		readerChan: make(chan readArg),
		readError:  make(chan error),
		writerChan: make(chan io.Writer),
		writeError: make(chan error),
	}
	ret.readDeadline.Store(time.Time{})
	ret.decoder.feeder = ret
	ret.writeDeadline.Store(time.Time{})
	ret.encoder.supportBinary = supportBinary
	ret.encoder.feeder = ret
	return ret
}

// FeedIn 数据流入
func (p *Payload) FeedIn(r io.Reader, supportBinary bool) error {
	select {
	case <-p.close:
		return p.load()
	default:
	}
	if !atomic.CompareAndSwapInt64(&p.feeding, 0, 1) {
		return primary.NewErr("", "read", false, primary.ErrOverlap)
	}
	defer atomic.StoreInt64(&p.feeding, 0)
	if ok := p.pauser.Working(); !ok {
		return primary.NewErr("", "payload", false, primary.ErrPaused)
	}
	defer p.pauser.Done()
	for {
		after, ok := p.readTimeout()
		if !ok {
			return p.Store("read", primary.ErrTimeout)
		}
		select {
		case <-p.close:
			return p.load()
		case <-after:
			continue
		case p.readerChan <- readArg{
			r:             r,
			supportBinary: supportBinary,
		}:
		}
		break
	}
	for {
		after, ok := p.readTimeout()
		if !ok {
			return p.Store("read", primary.ErrTimeout)
		}
		select {
		case <-after:
			continue
		case err := <-p.readError:
			return p.Store("read", err)
		}
	}
}

// FlushOut 刷新缓存
func (p *Payload) FlushOut(w io.Writer) error {
	select {
	case <-p.close:
		return p.load()
	default:
	}
	if !atomic.CompareAndSwapInt64(&p.flushing, 0, 1) {
		return primary.NewErr("", "write", false, primary.ErrOverlap)
	}
	defer atomic.StoreInt64(&p.flushing, 0)
	if ok := p.pauser.Working(); !ok {
		_, err := w.Write(p.encoder.NOOP())
		return err
	}
	defer p.pauser.Done()
	for {
		after, ok := p.writeTimeout()
		if !ok {
			return p.Store("write", primary.ErrTimeout)
		}
		select {
		case <-p.close:
			return p.load()
		case <-after:
			continue
		case <-p.pauser.PausingTrigger():
			_, err := w.Write(p.encoder.NOOP())
			return err
		case p.writerChan <- w:

		}
		break
	}
	for {
		after, ok := p.writeTimeout()
		if !ok {
			return p.Store("write", primary.ErrTimeout)
		}
		select {
		case <-after:
		case err := <-p.writeError:
			return p.Store("write", err)
		}
	}
}

// NextReader 下一个读取器
func (p *Payload) NextReader() (primary.FrameType, primary.PacketType, io.ReadCloser, error) {
	ft, pt, r, err := p.decoder.NextReader()
	return ft, pt, r, err
}

// SetReadDeadline 设置读取截止
func (p *Payload) SetReadDeadline(t time.Time) error {
	p.readDeadline.Store(t)
	return nil
}

// NextWriter 下一个写入器
func (p *Payload) NextWriter(ft primary.FrameType, pt primary.PacketType) (io.WriteCloser, error) {
	return p.encoder.NextWriter(ft, pt)
}

// SetWriteDeadline 设置写入截止
func (p *Payload) SetWriteDeadline(t time.Time) error {
	p.writeDeadline.Store(t)
	return nil
}

// Pause 暂停
func (p *Payload) Pause() {
	p.pauser.Pause()
}

// Resume 复位
func (p *Payload) Resume() {
	p.pauser.Resume()
}

// Close 关闭有效载荷
func (p *Payload) Close() error {
	p.closeOnce.Do(func() {
		close(p.close)
	})
	return nil
}

// Store 存储相关错误信息
func (p *Payload) Store(op string, err error) error {
	old := p.err.Load()
	if old == nil {
		if err == io.EOF || err == nil {
			return err
		}
		op := primary.NewErr("", op, false, err)
		p.err.Store(op)
		return op
	}
	return old.(error)
}

// getReader 获取读取器
func (p *Payload) getReader() (io.Reader, bool, error) {
	select {
	case <-p.close:
		return nil, false, p.load()
	default:
	}
	if ok := p.pauser.Working(); !ok {
		return nil, false, primary.NewErr("", "payload", false, primary.ErrPaused)
	}
	p.pauser.Done()
	for {
		after, ok := p.readTimeout()
		if !ok {
			return nil, false, p.Store("read", primary.ErrTimeout)
		}
		select {
		case <-p.close:
			return nil, false, p.load()
		case <-p.pauser.PausedTrigger():
			return nil, false, primary.NewErr("", "payload", false, primary.ErrPaused)
		case <-after:
			continue
		case arg := <-p.readerChan:
			return arg.r, arg.supportBinary, nil
		}
	}
}

// putReader 添加写入器
func (p *Payload) putReader(err error) error {
	select {
	case <-p.close:
		return p.load()
	default:
	}
	for {
		after, ok := p.readTimeout()
		if !ok {
			return p.Store("read", primary.ErrTimeout)
		}
		select {
		case <-p.close:
			return p.load()
		case <-after:
			continue
		case p.readError <- err:

		}
		return nil
	}
}

// readTimeout 读取超时
func (p *Payload) readTimeout() (<-chan time.Time, bool) {
	deadline := p.readDeadline.Load().(time.Time)
	wait := deadline.Sub(time.Now())
	if deadline.IsZero() {
		wait = math.MaxInt64
	}
	if wait <= 0 {
		return nil, false
	}
	return time.After(wait), true
}

func (p *Payload) getWriter() (io.Writer, error) {
	select {
	case <-p.close:
		return nil, p.load()
	default:
	}
	if ok := p.pauser.Working(); !ok {
		return nil, primary.NewErr("", "payload", false, primary.ErrPaused)
	}
	p.pauser.Done()
	for {
		after, ok := p.writeTimeout()
		if !ok {
			return nil, p.Store("write", primary.ErrTimeout)
		}
		select {
		case <-p.close:
			return nil, p.load()
		case <-p.pauser.PausedTrigger():
			return nil, primary.NewErr("", "payload", false, primary.ErrPaused)
		case <-after:
			continue
		case w := <-p.writerChan:
			return w, nil
		}
	}
}

// putWriter 添加写入器
func (p *Payload) putWriter(err error) error {
	select {
	case <-p.close:
		return p.load()
	default:
	}

	for {
		after, ok := p.writeTimeout()
		if !ok {
			return p.Store("write", primary.ErrTimeout)
		}
		ret := p.Store("write", err)
		select {
		case <-p.close:
			return p.load()
		case <-after:
			continue
		case p.writeError <- err:
			return ret
		}
	}
}

// writeTimeout 写入超时
func (p *Payload) writeTimeout() (<-chan time.Time, bool) {
	deadline := p.writeDeadline.Load().(time.Time)
	wait := deadline.Sub(time.Now())
	if deadline.IsZero() {
		wait = math.MaxInt64
	}
	if wait <= 0 {
		return nil, false
	}
	return time.After(wait), true
}

// load 加载
func (p *Payload) load() error {
	ret := p.err.Load()
	if ret == nil {
		return io.EOF
	}
	return ret.(error)
}
