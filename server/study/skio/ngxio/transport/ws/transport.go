package ws

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/znk_fullstack/server/study/skio/ngxio/base"
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

func (t *Transport) Dial(u *url.URL, requestHeader http.Header) (base.Conn, error) {
	dialer := ws.Dialer{}
}
