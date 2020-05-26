package polling

import (
	"net/http"
	"time"

	"github.com/znk_fullstack/server/study/skio/ngxio/base"
)

//Transport is the transport of polling
type Transport struct {
	Client      *http.Client
	CheckOrigin func(r *http.Request) bool
}

//Default is the default transport
var Default = &Transport{
	Client: &http.Client{
		Timeout: time.Minute,
	},
	CheckOrigin: nil,
}

//Name is the name of transport
func (t *Transport) Name() string {
	return "polling"
}

//Accept accepts a http request and create Conn.
func (t *Transport) Accept(w http.ResponseWriter, r *http.Request) (base.Conn, error) {

}
