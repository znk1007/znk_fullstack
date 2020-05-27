package ws

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/znk_fullstack/server/study/skio/ngxio/base"
	websocket "github.com/znk_fullstack/server/study/skio/ws"
)

//conn implements base.Conn
type conn struct {
	url          url.URL
	remoteHeader http.Header
	ws           wrapper
	closed       chan struct{}
	closeOnce    sync.Once
	base.FrameWriter
	base.FrameReader
}

func newConn(ws *websocket.Conn, url url.URL, header http.Header) base.Conn {

}
