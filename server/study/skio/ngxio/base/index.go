package base

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

//OpError is the error type usually returned by functions
//in the transport package
type OpError struct {
	URL string
	Op  string
	Err error
}

//OpErr new a *OpError
func OpErr(url, op string, err error) error {
	return &OpError{
		URL: url,
		Op:  op,
		Err: err,
	}
}

func (e *OpError) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Op, e.URL, e.Err.Error())
}

//Timeout returns true if the error is a timeout
func (e *OpError) Timeout() bool {
	if r, ok := e.Err.(net.Error); ok {
		return r.Timeout()
	}
	return false
}

//Temporary returns true if the error is temporary
func (e *OpError) Temporary() bool {
	if r, ok := e.Err.(net.Error); ok {
		return r.Temporary()
	}
	return false
}

//PacketType is the type of packet
type PacketType int

const (
	//OPEN is sent from the server when a new transport is opened (recheck).
	OPEN PacketType = iota
	//CLOSE is request the close of this transport but does not shutdown the
	//connection itself.
	CLOSE
	//PING is sent by client. Server should answer with a pong packet
	//containing the same data.
	PING
	//PONG is sent by the server to respond to ping packets.
	PONG
	//MESSAGE is actual message, client and server should call their callbacks
	//with the data.
	MESSAGE
	//UPGRADE is sent before ngxio switches a transport to test if server
	//and client can communicate over this transport. If this test succeed,
	//when client sends an upgrade packets which requests the server to flush
	//tis cache on the old transport and switch to the new transport.
	UPGRADE
	//NOOP is a noop packet. Used primarily to force a poll cycle when an
	//incoming websocket connection is received.
	NOOP
)

func (pt PacketType) String() string {
	switch pt {
	case OPEN:
		return "open"
	case CLOSE:
		return "close"
	case PING:
		return "ping"
	case PONG:
		return "pong"
	case MESSAGE:
		return "message"
	case UPGRADE:
		return "upgrade"
	case NOOP:
		return "noop"
	}
	return "unknown"
}

//StringByte converts a PacketType to byte in string
func (pt PacketType) StringByte() byte {
	return byte(pt) + '0'
}

//BinaryByte converts a PacketType to byte in binary.
func (pt PacketType) BinaryByte() byte {
	return byte(pt)
}

//ByteToPacketType converts a byte to PacketType.
func ByteToPacketType(b byte, ft FrameType) PacketType {
	if ft == FrameString {
		b -= '0'
	}
	return PacketType(b)
}

// FrameType is the type of frames.
type FrameType byte

const (
	//FrameString identifies a string frame.
	FrameString FrameType = iota
	//FrameBinary identifies a binary frame.
	FrameBinary
)

//ByteToFrameType converts a byte to FrameType.
func ByteToFrameType(b byte) FrameType {
	return FrameType(b)
}

//Byte return type in byte
func (t FrameType) Byte() byte {
	return byte(t)
}

//FrameReader reads a frame. It need be closed before next reading.
type FrameReader interface {
	NextReader() (FrameType, PacketType, error)
}

//FrameWriter writes a frame. It need be closed before next writing.
type FrameWriter interface {
	NextWriter(ft FrameType, pt PacketType) (io.WriteCloser, error)
}

var chars = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_")

//Timestamp returns a string based on different nano time.
func Timestamp() string {
	now := time.Now().UnixNano()
	ret := make([]byte, 0, 16)
	for now > 0 {
		ret = append(ret, chars[int(now%int64(len(chars)))])
		now = now / int64(len(chars))
	}
	return string(ret)
}

//Conn connection interface
type Conn interface {
	FrameReader
	FrameWriter
	io.Closer
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}

//ConnParamters is connection parameter of server
type ConnParamters struct {
	PingInterval time.Duration
	PingTimeout  time.Duration
	SID          string
	Upgrades     []string
}

type jsonParameters struct {
	SID          string   `json:"sid"`
	Upgrades     []string `json:"upgrades"`
	PingInterval int      `json:"pingInterval"`
	PingTimeout  int      `json:"pingTimeout"`
}

//ReadConnParameters reads ConnParameters from r.
func ReadConnParameters(r io.Reader) (ConnParamters, error) {
	var param jsonParameters
	if err := json.NewDecoder(r).Decode(&param); err != nil {
		return ConnParamters{}, err
	}
	return ConnParamters{
		SID:          param.SID,
		Upgrades:     param.Upgrades,
		PingInterval: time.Duration(param.PingInterval) * time.Millisecond,
		PingTimeout:  time.Duration(param.PingTimeout) * time.Microsecond,
	}, nil
}

//WriteTo writes to w with json format.
func (cp ConnParamters) WriteTo(w io.Writer) (int64, error) {
	arg := jsonParameters{
		SID:          cp.SID,
		Upgrades:     cp.Upgrades,
		PingInterval: int(cp.PingInterval / time.Microsecond),
		PingTimeout:  int(cp.PingTimeout / time.Millisecond),
	}
	writer := writer{
		w: w,
	}
	err := json.NewEncoder(&writer).Encode(arg)
	return writer.i, err
}

type writer struct {
	i int64
	w io.Writer
}

func (w *writer) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	w.i += int64(n)
	return n, err
}
