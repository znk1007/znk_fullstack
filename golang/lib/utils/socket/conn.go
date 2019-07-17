package socket

import (
	sysIO "io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gogo/protobuf/io"

	socket "znk/golang/lib/utils/socket/protos/generated"
)

// Conn 连接接口
type Conn interface {
	FrameReader
	FrameWriter
	sysIO.Closer
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

// ReadConnParams 读取连接参数
func ReadConnParams(r sysIO.Reader) (socket.ConnParams, error) {
	param := socket.ConnParams{}
	reader := io.NewFullReader(r, 1024)
	defer reader.Close()
	err := reader.ReadMsg(&param)
	if err != nil {
		return socket.ConnParams{}, err
	}
	return param, nil
}

// WriteTo 写入数据
func WriteTo(params socket.ConnParams, w sysIO.Writer) (int, error) {
	writer := io.NewFullWriter(w)
	defer writer.Close()
	return params.Size(), writer.WriteMsg(&params)
}
