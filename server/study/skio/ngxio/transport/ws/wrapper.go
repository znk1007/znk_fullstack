package ws

import (
	"net/http"
	"net/url"
)

type wrapper struct {
	url          url.URL
	remoteHeader http.Header
}
