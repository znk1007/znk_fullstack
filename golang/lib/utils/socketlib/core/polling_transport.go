package core

import (
	"net/http"
	"net/url"
	"time"
)

//PollingTransport 轮询传输
type PollingTransport struct {
	Client *http.Client
}

// Default 默认
var Default = &PollingTransport{
	Client: &http.Client{
		Timeout: time.Minute,
	},
}

// Name 名称
func (pt *PollingTransport) Name() string {
	return "polling"
}

// Accept 接收
func (pt *PollingTransport) Accept(w http.ResponseWriter, r *http.Request) (Conn, error) {
	conn := newPollingServer(r)
	return conn, nil
}

// Dial 拨号
func (pt *PollingTransport) Dial(u *url.URL, requestHeader http.Header) (Conn, error) {
	query := u.Query()
	query.Set("transport", pt.Name())
	u.RawQuery = query.Encode()
	return dial(pt.Client, u, requestHeader)
}
