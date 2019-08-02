package core

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

// writeManager 写入管理
type writeManager interface {
	getWriter() (io.Writer, error)
	addWriter(error) error
}

// readManager 读取管理
type readManager interface {
	getReader() (io.Reader, bool, error)
	addReader(error) error
}

// byteReader 字节读取器
type byteReader interface {
	ReadByte() (byte, error)
	io.Reader
}

// dataReader 数据读取器
type dataReader interface {
	NextReader() (pbs.DataType, pbs.PacketType, io.ReadCloser, error)
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
