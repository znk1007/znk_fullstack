package core

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

// dataReader 数据读取器
type dataReader interface {
	NextReader() (pbs.DataType, io.ReadCloser, error)
}

// dataWriter 数据写入器
type dataWriter interface {
	NextWriter(dt pbs.DataType, pt pbs.PacketType) (io.WriteCloser, error)
}

// Conn 连接接口
type Conn interface {
	dataReader
	dataWriter
	io.Closer
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
	ID() string
	Context() interface{}
	SetContext(v interface{})
	Namespace() string
	Emit(msg string, v ...interface{})
}
