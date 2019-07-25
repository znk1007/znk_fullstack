package websocket

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/znk_fullstack/golang/lib/utils/socket/packet"

	"github.com/znk_fullstack/golang/lib/utils/socket/transport"

	"github.com/gorilla/websocket"
	"github.com/znk_fullstack/golang/lib/utils/socket/primary"
)

type wrapper struct {
	*websocket.Conn
}

// newWrapper 创建包装
func newWrapper(conn *websocket.Conn) wrapper {
	return wrapper{
		Conn: conn,
	}
}

func (w wrapper) NextReader() (primary.FrameType, io.ReadCloser, error) {
	t, r, e := w.Conn.NextReader()
	if e != nil {
		return 0, nil, e
	}
	ret := ioutil.NopCloser(r)
	switch t {
	case websocket.TextMessage:
		return primary.FrameString, ret, nil
	case websocket.BinaryMessage:
		return primary.FrameBinary, ret, nil
	}
	return 0, nil, transport.ErrInvalidFrame
}

// NextWriter 下一个写入器
func (w wrapper) NextWriter(t primary.FrameType) (io.WriteCloser, error) {
	var target int
	switch t {
	case primary.FrameString:
		target = websocket.TextMessage
	case primary.FrameBinary:
		target = websocket.BinaryMessage
	default:
		return nil, transport.ErrInvalidFrame
	}
	return w.Conn.NextWriter(target)
}

type conn struct {
	url          url.URL
	remoteHeader http.Header
	wp           wrapper
	closed       chan struct{}
	closeOnce    sync.Once
	primary.FrameWriter
	primary.FrameReader
}

func newConn(ws *websocket.Conn, url url.URL, header http.Header) primary.Conn {
	w := newWrapper(ws)
	closed := make(chan struct{})
	return &conn{
		url:          url,
		remoteHeader: header,
		wp:           w,
		closed:       closed,
		FrameReader:  packet.NewDecoder(w),
		FrameWriter:  packet.NewEncoder(w),
	}
}

func (c *conn) URL() url.URL {
	return c.url
}

func (c *conn) RemoteHeader() http.Header {
	return c.remoteHeader
}

func (c *conn) LocalAddr() net.Addr {
	return c.wp.LocalAddr()
}

func (c *conn) RemoteAddr() net.Addr {
	return c.wp.RemoteAddr()
}

func (c *conn) SetReadDeadline(t time.Time) error {
	return c.wp.SetReadDeadline(t)
}

func (c *conn) SetWriteDeadline(t time.Time) error {
	return c.wp.SetWriteDeadline(t)
}

func (c *conn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	<-c.closed
}

func (c *conn) Close() error {
	c.closeOnce.Do(func() {
		close(c.closed)
	})
	return c.wp.Close()
}

// DialError 拨号错误
type DialError struct {
	error
	Response *http.Response
}

// Transport 传输实现
type Transport struct {
	ReadBufferSize   int
	WriteBufferSize  int
	NetDial          func(network, addr string) (net.Conn, error)
	Proxy            func(*http.Request) (*url.URL, error)
	TLSClientConfig  *tls.Config
	HandshakeTimeout time.Duration
	Subprotocols     []string
	CheckOrigin      func(r *http.Request) bool
}

// Default 默认传输
var Default = &Transport{}

// Name 传输名称
func (t *Transport) Name() string {
	return "websocket"
}

// Dial 传输拨号
func (t *Transport) Dial(u *url.URL, requestHeader http.Header) (primary.Conn, error) {
	dialer := websocket.Dialer{
		ReadBufferSize:   t.ReadBufferSize,
		WriteBufferSize:  t.WriteBufferSize,
		NetDial:          t.NetDial,
		Proxy:            t.Proxy,
		TLSClientConfig:  t.TLSClientConfig,
		HandshakeTimeout: t.HandshakeTimeout,
		Subprotocols:     t.Subprotocols,
	}
	switch u.Scheme {
	case "http":
		u.Scheme = "ws"
	case "https":
		u.Scheme = "wss"
	}
	query := u.Query()
	query.Set("transport", t.Name())
	query.Set("t", primary.NewSocketID().String())
	u.RawQuery = query.Encode()
	c, resp, err := dialer.Dial(u.String(), requestHeader)
	if err != nil {
		return nil, DialError{
			error:    err,
			Response: resp,
		}
	}
	return newConn(c, *u, resp.Header), nil
}

// Accept 接收数据
func (t *Transport) Accept(w http.ResponseWriter, r *http.Request) (primary.Conn, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  t.ReadBufferSize,
		WriteBufferSize: t.WriteBufferSize,
		CheckOrigin:     t.CheckOrigin,
	}
	c, err := upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		return nil, err
	}
	return newConn(c, *r.URL, r.Header), nil
}
