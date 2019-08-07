package core

import (
	"io"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/znk_fullstack/golang/lib/utils/socketlib/protos/pbs"
)

type readArg struct {
	r             io.Reader
	supportBinary bool
}

type payload struct {
	close     chan struct{}
	closeOnce sync.Once
	err       atomic.Value

	pp *payloadPauser

	readerChan   chan readArg
	feeding      int64
	readError    chan error
	readDeadline atomic.Value
	pd           payloadDecoder

	writerChan    chan io.Writer
	flushing      int64
	writeError    chan error
	writeDeadline atomic.Value
	pe            payloadEncoder
}

// NewPayload 创建payload
func NewPayload(supportBinary bool) *payload {
	ret := &payload{
		close:      make(chan struct{}),
		pp:         newPayloadPauser(),
		readerChan: make(chan readArg),
		readError:  make(chan error),
		writerChan: make(chan io.Writer),
		writeError: make(chan error),
	}
	ret.readDeadline.Store(time.Time{})
	ret.pd.readM = ret
	ret.writeDeadline.Store(time.Time{})
	ret.pe.supportBinary = supportBinary
	ret.pe.writeM = ret
	return ret
}

func (p *payload) FeedIn(r io.Reader, supportBinary bool) error {
	select {
	case <-p.close:
		return p.load()
	default:
	}
	if !atomic.CompareAndSwapInt64(&p.feeding, 0, 1) {
		return newErr("", "read", errOverlap)
	}
	defer atomic.StoreInt64(&p.feeding, 0)
	if ok := p.pp.Working(); !ok {
		return newErr("", "payload", errPaused)
	}
	defer p.pp.Done()
	for {
		after, ok := p.readTimeout()
		if !ok {
			return p.Store("read", errTimeout)
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
			return p.Store("read", errTimeout)
		}
		select {
		case <-after:
			continue
		case err := <-p.readError:
			return p.Store("read", err)
		}
	}
}

func (p *payload) FlushOut(w io.Writer) error {
	select {
	case <-p.close:
		return p.load()
	default:
	}
	if !atomic.CompareAndSwapInt64(&p.flushing, 0, 1) {
		return newErr("", "write", errOverlap)
	}
	defer atomic.StoreInt64(&p.flushing, 0)
	if ok := p.pp.Working(); !ok {
		_, err := w.Write(p.pe.Noop())
		return err
	}
	defer p.pp.Done()
	for {
		after, ok := p.writeTimeout()
		if !ok {
			return p.Store("write", errTimeout)
		}
		select {
		case <-p.close:
			return p.load()
		case <-after:
			continue
		case <-p.pp.PausingTrigger():
			_, err := w.Write(p.pe.Noop())
			return err
		case p.writerChan <- w:
		}
		break
	}
	for {
		after, ok := p.writeTimeout()
		if !ok {
			return p.Store("write", errTimeout)
		}
		select {
		case <-after:
		case err := <-p.writeError:
			return p.Store("write", err)
		}
	}
}

func (p *payload) NextReader() (pbs.DataType, pbs.PacketType, io.ReadCloser, error) {
	ft, pt, r, err := p.pd.NextReader()
	return ft, pt, r, err
}

func (p *payload) SetReadDeadline(t time.Time) error {
	p.readDeadline.Store(t)
	return nil
}

func (p *payload) NextWriter(dt pbs.DataType, pt pbs.PacketType) (io.WriteCloser, error) {
	return p.pe.NextWriter(dt, pt)
}

func (p *payload) SetWriteDeadline(t time.Time) error {
	p.writeDeadline.Store(t)
	return nil
}

func (p *payload) Pause() {
	p.pp.Pause()
}

func (p *payload) Resume() {
	p.pp.Resume()
}

func (p *payload) Close() error {
	p.closeOnce.Do(func() {
		close(p.close)
	})
	return nil
}

func (p *payload) Store(op string, err error) error {
	old := p.err.Load()
	if old == nil {
		if err == io.EOF || err == nil {
			return err
		}
		op := newErr("", op, err)
		p.err.Store(op)
		return op
	}
	return old.(error)
}

func (p *payload) getReader() (io.Reader, bool, error) {
	select {
	case <-p.close:
		return nil, false, p.load()
	default:
	}
	if ok := p.pp.Working(); !ok {
		return nil, false, newErr("", "payload", errPaused)
	}
	p.pp.Done()
	for {
		after, ok := p.readTimeout()
		if !ok {
			return nil, false, p.Store("read", errTimeout)
		}
		select {
		case <-p.close:
			return nil, false, p.load()
		case <-p.pp.PausedTrigger():
			return nil, false, newErr("", "payload", errPaused)
		case <-after:
			continue
		case arg := <-p.readerChan:
			return arg.r, arg.supportBinary, nil
		}
	}
}

func (p *payload) addReader(err error) error {
	select {
	case <-p.close:
		return p.load()
	default:
	}
	for {
		after, ok := p.readTimeout()
		if !ok {
			return p.Store("read", errTimeout)
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

func (p *payload) readTimeout() (<-chan time.Time, bool) {
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

func (p *payload) getWriter() (io.Writer, error) {
	select {
	case <-p.close:
		return nil, p.load()
	default:
	}
	if ok := p.pp.Working(); !ok {
		return nil, newErr("", "payload", errPaused)
	}
	p.pp.Done()
	for {
		after, ok := p.writeTimeout()
		if !ok {
			return nil, p.Store("write", errTimeout)
		}
		select {
		case <-p.close:
			return nil, p.load()
		case <-p.pp.PausedTrigger():
			return nil, newErr("", "payload", errPaused)
		case <-after:
			continue
		case w := <-p.writerChan:
			return w, nil
		}
	}
}
func (p *payload) addWriter(err error) error {
	select {
	case <-p.close:
		return p.load()
	default:
	}
	for {
		after, ok := p.writeTimeout()
		if !ok {
			return p.Store("write", errTimeout)
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

func (p *payload) writeTimeout() (<-chan time.Time, bool) {
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

func (p *payload) load() error {
	ret := p.err.Load()
	if ret == nil {
		return io.EOF
	}
	return ret.(error)
}
