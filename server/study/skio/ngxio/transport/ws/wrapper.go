package ws

import (
	"net/http"
	"net/url"
	"sync"

	websocket "github.com/znk_fullstack/server/study/skio/ws"
)

type wrapper struct {
	url          url.URL
	remoteHeader http.Header
	readLocker   *sync.Mutex
}

func newWrapper(conn *websocket.Conn) wrapper {

}
