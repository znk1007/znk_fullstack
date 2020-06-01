package sio

import (
	"net"
	"net/http"
	"net/url"

	"github.com/znk_fullstack/server/study/skio/ngxio"
	"github.com/znk_fullstack/server/study/skio/sio/parser"
)

//Conn is a connection in sio
type Conn interface {
	//ID returns session id
	ID() string
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	//Context of this connection. You can save one context for one
	//connection, and share it between all handlers. The handlers
	//is called in one goroutine, so no need to lock context if it
	//only be accessed in one connection.
	Context() interface{}
	SetContext(v interface{})
	Namespace() string
	Emit(msg string, v ...interface{})

	//Broadcast server side apis
	Join(room string)
	Leave(room string)
	LeaveAll()
	Rooms() []string
}

type errorMessage struct {
	namespace string
	error
}

type writePacket struct {
	header parser.Header
	data   []interface{}
}

type conn struct {
	ngxio.Conn
	encoder   *parser.Encoder
	decoder   *parser.Decoder
	errorChan chan errorMessage
	writeChan chan writePacket
	quitChan  chan struct{}
	handlers map[string]
}
