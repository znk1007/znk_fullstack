package polling

import (
	"errors"
	"mime"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"github.com/znk_fullstack/golang/lib/utils/socket/payload"
)

func mimeSupportBinary(m string) (bool, error) {
	t, p, e := mime.ParseMediaType(m)
	if e != nil {
		return false, e
	}
	switch t {
	case "application/octet-stream":
		return true, nil
	case "text/plain":
		charset := strings.ToLower(p["charset"])
		if charset != "utf-8" {
			return false, errors.New("invalid charset")
		}
		return false, nil
	}
	return false, errors.New("invalid content-type")
}

// Addr 网络请求地址
type Addr struct {
	Host string
}

// Network 网络连接类型
func (a Addr) Network() string {
	return "tcp"
}

// String 端口
func (a Addr) String() string {
	return a.Host
}

// Transport 轮询传输
type Transport struct {
	Client *http.Client
}

// Default 默认传输
var Default = &Transport{
	Client: &http.Client{
		Timeout: time.Minute,
	},
}

// Name 传输名称
func (t *Transport) Name() string {
	return "polling"
}

// func (t *Transport) Accept(w http.ResponseWriter, r *http.Request) (primary.Conn, error) {
// 	conn :=
// }

type clientConn struct {
	*payload.Payload
	httpClient   *http.Client
	request      http.Request
	remoteHeader atomic.Value
}

func dial(client *http.Client, url *url.URL, requestHeader http.Header) (*clientConn, error) {
	if client == nil {
		client = &http.Client{}
	}
	req, err := http.NewRequest("", url.String(), nil)
	if err != nil {
		return nil, err
	}
	for k, v := range requestHeader {
		req.Header[k] = v
	}
	supportBinary := req.URL.Query().Get("base64") == ""
	if supportBinary {
		req.Header.Set("Content-Type", "application/octet-stream")
	} else {
		req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	}
	ret := &clientConn{
		Payload:    payload.New(supportBinary),
		httpClient: client,
		request:    *req,
	}
}
