package ngxio

import (
	"fmt"
	"io"
	"net"
)

//opError is the error type usually returned by functions
//in the transport package
type opError struct {
	URL string
	Op  string
	Err error
}

//opErr new a *opError
func opErr(url, op string, err error) error {
	return &opError{
		URL: url,
		Op:  op,
		Err: err,
	}
}

func (e *opError) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Op, e.URL, e.Err.Error())
}

//Timeout returns true if the error is a timeout
func (e *opError) Timeout() bool {
	if r, ok := e.Err.(net.Error); ok {
		return r.Timeout()
	}
	return false
}

//Temporary returns true if the error is temporary
func (e *opError) Temporary() bool {
	if r, ok := e.Err.(net.Error); ok {
		return r.Temporary()
	}
	return false
}

type packetType int

// frameType is the type of frames.
type frameType byte

const (
	//FrameString identifies a string frame.
	frameString frameType = iota
	//FrameBinary identifies a binary frame.
	frameBinary
)

//byteToFrameType converts a byte to frameType.
func byteToFrameType(b byte) frameType {
	return frameType(b)
}

//byte return type in byte
func (t frameType) byte() byte {
	return byte(t)
}

//frameReader frameReader reads a frame. It need be closed before next reading.
type frameReader interface {
	nextReader() (frameType, packetType, error)
}

//frameWriter writes a frame. It need be closed before next writing.
type frameWriter interface {
	nextWriter(ft frameType, pt packetType) (io.WriteCloser, error)
}
