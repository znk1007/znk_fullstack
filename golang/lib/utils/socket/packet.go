package socket

import "io"

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
