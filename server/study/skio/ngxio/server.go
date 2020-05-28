package ngxio

import (
	"net/http"
	"time"

	"github.com/znk_fullstack/server/study/skio/ngxio/transport"
)

func defaultChecker(*http.Request) (http.Header, error) {
	return nil, nil
}

func defaultInitor(*http.Request, Conn) {}

//Options is options to create a server.
type Options struct {
	RequestChecker func(*http.Request) (http.Header, error)
	ConnInitor     func(*http.Request, Conn)
	PingTimeout    time.Duration
	PingInterval   time.Duration
	Transports     []transport.Transport
	SessionIDGenerator
}
