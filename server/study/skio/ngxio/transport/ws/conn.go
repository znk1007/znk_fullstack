package ws

import (
	"net/http"
	"net/url"
)

//conn implements base.Conn
type conn struct {
	url          url.URL
	remoteHeader http.Header
}
