package primary

import (
	"errors"
	"fmt"
	sysIO "io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gogo/protobuf/io"
	socket "github.com/znk_fullstack/golang/lib/utils/socket/protos/generated"
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

// Error 错误接口
type Error interface {
	Error() string
	Temporary() bool
}

// Errs socket相关错误
type Errs struct {
	URL       string
	Operation string
	Err       error
	IsNet     bool
}

// NewErr 创建错误
func NewErr(url, operation string, isNet bool, err error) error {
	return &Errs{
		URL:       url,
		Operation: operation,
		Err:       err,
	}
}

func (e *Errs) Error() string {
	if e.URL == "" {
		return fmt.Sprintf("%s:%s", e.Operation, e.Err.Error())
	}
	return fmt.Sprintf("%s %s:%s", e.Operation, e.URL, e.Err.Error())
}

// Timeout 超时错误
func (e *Errs) Timeout() bool {
	if e.IsNet {
		if r, ok := e.Err.(net.Error); ok {
			return r.Timeout()
		}
	}
	return false
}

// Temporary 临时错误
func (e *Errs) Temporary() bool {
	if e.IsNet {
		if err, ok := e.Err.(net.Error); ok {
			return err.Temporary()
		}
	}
	if err, ok := e.Err.(Error); ok {
		return err.Temporary()
	}
	return false
}

// RetryError 重试错误
type RetryError struct {
	err string
}

func (e RetryError) Error() string {
	return e.err
}

// Temporary 重试临时错误
func (e RetryError) Temporary() bool {
	return true
}

var (
	// ErrPaused 暂停错误
	ErrPaused = RetryError{"paused"}
	// ErrTimeout 超时
	ErrTimeout = RetryError{"timeout"}
	// ErrInvalidPayload 无效负载
	ErrInvalidPayload = errors.New("invalid payload")
	// ErrDrain 无效输出
	ErrDrain = errors.New("drain")
	// ErrOverlap 重叠错误
	ErrOverlap = errors.New("overlap")
)

// PacketType 打包类型
type PacketType int

const (
	// Open 开启状态，当传输打开时，服务器处于该状态
	Open PacketType = iota
	// Close 关闭当前传输，但并未断开连接
	Close
	// Ping 客户端发送的消息，服务器将返回pong包，确保正在连接
	Ping
	// Pong 服务器响应客户端Ping而发送的包
	Pong
	// Message 传输内容，服务端，客户端都需回调的内容
	Message
	// Upgrade 更新传输缓存，旧传输更新至新传输
	Upgrade
	// Noop 连接websocket时，强制轮询的包
	Noop
)

// ToString 转字符串
func (t PacketType) ToString() string {
	switch t {
	case Open:
		return "open"
	case Close:
		return "close"
	case Ping:
		return "ping"
	case Pong:
		return "pong"
	case Message:
		return "message"
	case Upgrade:
		return "upgrade"
	case Noop:
		return "noop"

	}
	return "unkown"
}

// ToStringByte 字符位
func (t PacketType) ToStringByte() byte {
	return byte(t) + '0'
}

// ToBinaryByte 二进制位
func (t PacketType) ToBinaryByte() byte {
	return byte(t)
}

// ToPacketType 转包类型
func ToPacketType(b byte, t FrameType) PacketType {
	if t == FrameString {
		b -= '0'
	}
	return PacketType(b)
}

// FrameType 数据帧
type FrameType byte

const (
	// FrameString 字符串数据帧
	FrameString FrameType = iota
	// FrameBinary 二进制数据帧
	FrameBinary
)

// ToFrameType 转帧类型
func ToFrameType(b byte) FrameType {
	return FrameType(b)
}

// Byte 帧的二进制
func (t FrameType) Byte() byte {
	return byte(t)
}

// FrameReader 读取数据帧
type FrameReader interface {
	NextReader() (FrameType, PacketType, io.ReadCloser, error)
}

// FrameWriter 写入数据帧
type FrameWriter interface {
	NextWriter(ft FrameType, pt PacketType) (io.WriteCloser, error)
}
