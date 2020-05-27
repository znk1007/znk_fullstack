package ws

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/znk_fullstack/server/study/skio/ngxio/base"
	websocket "github.com/znk_fullstack/server/study/skio/ws"
)

//DialError is the error when dialing to a server.
//It saves Response from server.
type DialError struct {
	error
	Response *http.Response
}

//Transport is websocket transport
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

//Default is default transport.
var Default = &Transport{}

//Name is the name of websocket transport
func (t *Transport) Name() string {
	return "websocket"
}

//Dial creates a new client connection.
func (t *Transport) Dial(u *url.URL, requestHeader http.Header) (base.Conn, error) {
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
	query.Set("t", base.Timestamp())
	u.RawQuery = query.Encode()
	c, resp, err := dialer.Dial(u.String(), requestHeader)
	if err != nil {
		return nil, DialError{
			error:    err,
			Response: resp,
		}
	}
	return n
}
