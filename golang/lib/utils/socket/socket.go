package socket

import (
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	protos "github.com/znk_fullstack/golang/lib/utils/socket/protos/generated"

	"github.com/znk_fullstack/golang/lib/utils/socket/transport"

	"github.com/znk_fullstack/golang/lib/utils/socket/primary"
)

type writer struct {
	io.WriteCloser
	locker    *sync.RWMutex
	closeOnce sync.Once
}

func newWriter(w io.WriteCloser, locker *sync.RWMutex) *writer {
	return &writer{
		WriteCloser: w,
		locker:      locker,
	}
}

func (w *writer) Close(err error) {
	w.closeOnce.Do(func() {
		w.locker.Lock()
		defer w.locker.Unlock()
		err = w.WriteCloser.Close()
	})
}

func (w *writer) Write(bs []byte) (int, error) {
	w.locker.Lock()
	defer w.locker.Unlock()
	return w.WriteCloser.Write(bs)
}

type reader struct {
	io.ReadCloser
	locker    *sync.RWMutex
	closeOnce sync.Once
}

func newReader(r io.ReadCloser, locker *sync.RWMutex) *reader {
	return &reader{
		ReadCloser: r,
		locker:     locker,
	}
}

func (r *reader) Close() (err error) {
	r.closeOnce.Do(func() {
		r.locker.Lock()
		io.Copy(ioutil.Discard, r.ReadCloser)
		err = r.ReadCloser.Close()
		r.locker.Unlock()
	})
	return
}

// FrameType 数据类型
type FrameType primary.FrameType

const (
	// Text 文本类型数据
	Text = FrameType(primary.FrameString)
	// Binary 二进制类型数据
	Binary = FrameType(primary.FrameBinary)
)

// Conn 连接接口
type Conn interface {
	ID() string
	NextReader() (FrameType, io.ReadCloser, error)
	NextWriter(ft FrameType) (io.WriteCloser, error)
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	SetContext(v interface{})
	Context() interface{}
}

// Dialer 拨号器
type Dialer struct {
	Transports []transport.Transport
}

// Dial 拨号
func (d *Dialer) Dial(urlStr string, requestHeader http.Header) (Conn, error) {
	u, e := url.Parse(urlStr)
	if e != nil {
		return nil, e
	}
	query := u.Query()
	query.Set("EIO", "3")
	u.RawQuery = query.Encode()
	var conn primary.Conn
	for idx := len(d.Transports) - 1; idx >= 0; idx-- {
		if conn != nil {
			conn.Close()
		}
		t := d.Transports[idx]
		conn, e := t.Dial(u, requestHeader)
		if e != nil {
			continue
		}
		var params protos.ConnParams
		if p, ok := conn.(transport.Opener); ok {
			params, e = p.Open()
			if e != nil {
				continue
			}
		} else {
			var pt primary.PacketType
			var r io.ReadCloser
			_, pt, r, e = conn.NextReader()
			if e != nil {
				continue
			}
			func() {
				defer r.Close()
				if pt != primary.Open {
					e = errors.New("invalid open")
					return
				}
				params, e = primary.ReadConnParams(r)
				if e != nil {
					return
				}
			}()
		}
		if e != nil {
			continue
		}
		ret := &client{
			conn:      conn,
			params:    params,
			transport: t.Name(),
			close:     make(chan struct{}),
		}
		go ret.server()
		return ret, nil
	}
	return nil, e
}

type client struct {
	conn      primary.Conn
	params    protos.ConnParams
	transport string
	context   interface{}
	close     chan struct{}
	closeOnce sync.Once
}

func (c *client) SetContext(v interface{}) {
	c.context = v
}

func (c *client) Context() interface{} {
	return c.context
}

func (c *client) ID() string {
	return c.params.SID
}

func (c *client) Transport() string {
	return c.transport
}

func (c *client) Close() error {
	c.closeOnce.Do(func() {
		close(c.close)
	})
	return c.conn.Close()
}

func (c *client) NextReader() (FrameType, io.ReadCloser, error) {
	for {
		ft, pt, r, err := c.conn.NextReader()
		if err != nil {
			return 0, nil, err
		}
		switch pt {
		case primary.Pong:
			c.conn.SetReadDeadline(time.Now().Add(time.Duration(c.params.PingTimeout)))
		case primary.Close:
			c.Close()
			return 0, nil, io.EOF
		case primary.Message:
			return FrameType(ft), r, nil
		}
		r.Close()
	}
}

func (c *client) NextWriter(ft FrameType) (io.WriteCloser, error) {
	return c.conn.NextWriter(primary.FrameType(ft), primary.Message)
}

func (c *client) URL() url.URL {
	return c.conn.URL()
}

func (c *client) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *client) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *client) RemoteHeader() http.Header {
	return c.conn.RemoteHeader()
}

func (c *client) server() {
	defer c.conn.Close()
	for {
		select {
		case <-c.close:
			return
		case <-time.After(time.Duration(c.params.PingInterval)):
		}
		w, err := c.conn.NextWriter(primary.FrameString, primary.Ping)
		if err != nil {
			return
		}
		if err := w.Close(); err != nil {
			return
		}
		c.conn.SetWriteDeadline(time.Now().Add(time.Duration(c.params.PingTimeout)))
	}
}
