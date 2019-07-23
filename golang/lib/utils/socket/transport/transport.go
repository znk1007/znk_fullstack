package transport

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/znk_fullstack/golang/lib/utils/socket/primary"
	socket "github.com/znk_fullstack/golang/lib/utils/socket/protos/generated"
)

// Checker 校验请求
type Checker func(*http.Request) (http.Header, error)

// ErrInvalidFrame 无效帧率
var ErrInvalidFrame = errors.New("invalid frame type")

// ErrInvalidPacket 无效包
var ErrInvalidPacket = errors.New("invalid packet type")

// HTTPError HTTP请求错误
type HTTPError interface {
	Code() int
}

// Transport 传输接口
type Transport interface {
	Name() string
	Accept(w http.ResponseWriter, r *http.Request) (primary.Conn, error)
	Dial(u *url.URL, requestHeader http.Header) (primary.Conn, error)
}

// Pauser 暂停接口
type Pauser interface {
	Pasue()
	Resume()
}

// Opener 打开连接
type Opener interface {
	Open() (socket.ConnParams, error)
}

// Manager 传输管理
type Manager struct {
	order     []string
	transport map[string]Transport
}

// NewManager 创建传输管理
func NewManager(transports []Transport) *Manager {
	tranMap := make(map[string]Transport)
	names := make([]string, len(transports))
	for i, t := range transports {
		names[i] = t.Name()
		tranMap[t.Name()] = t
	}
	return &Manager{
		order:     names,
		transport: tranMap,
	}
}
