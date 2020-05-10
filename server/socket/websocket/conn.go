package ws

import (
	"errors"
	"io"
	"net"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
)

const (
	//Frame header byte 0 bits from Section 5.2 of RFC 6455
	finalBit = 1 << 7
	rsv1Bit  = 1 << 6
	rsv2Bit  = 1 << 5
	rsv3Bit  = 1 << 4
	//Frame header byte 1 bits from Section 5.2 of RFC 6455
	maskBit = 1 << 7
	//Fixed header + length +mask
	maxFrameHeaderSize         = 2 + 8 + 4
	maxControlFramePayloadSize = 125

	writeWait = time.Second

	defaultReadBufferSize  = 4096
	defaultWriteBufferSize = 4096

	continuationFrame = 0
	noFrame           = -1
)

//Close codes Defined in RFC 6455, section 11.7
const (
	CloseNormalClosure           = 1000
	CloseGoingAway               = 1001
	CloseProtocolError           = 1002
	CloseUnsupportedData         = 1003
	CloseNoStatusReceived        = 1005
	CloseAbnormalClosure         = 1006
	CloseInvalidFramePayloadData = 1007
	ClosePolicyViolation         = 1008
	CloseMessageTooBig           = 1009
	CloseMandatoryExtension      = 1010
	CloseInternalServerErr       = 1011
	CloseServiceRestart          = 1012
	CloseTryAgainLater           = 1013
	CloseTLSHandshake            = 1015
)

//The message typs are defined in RFC 6455, section 11.8
const (
	//TextMessage denotes a text data message. The text message payload is
	//interpreted as UTF-8 encoded text data.
	TextMessage = 1
	//BinaryMessage denotes a binary data message
	BinaryMessage = 2
	//CloseMessage denotes a close control message. The optoinal message
	//payload contains a numeric code and text. Use the FormatCloseMessage
	//function to format a close message payload
	CloseMessage = 8
	//PingMessage denotes a ping control message. The optional message payload
	//is UTF-8 encoded text.
	PingMessage = 9
	//PongMessage denotes a ping control message. The optional message payload
	//is UTF-8 encoded text.
	PongMessage = 10
)

//ErrCloseSent is returned when the application writes a message to the
//connection after sending a close message
var ErrCloseSent = errors.New("ws: close sent")

//ErrReadLimit is returned when reading a message that is larget that the
//read limit set for the connection.
var ErrReadLimit = errors.New("ws: read limit exceeded")

type netError struct {
	msg       string
	temporary bool
	timeout   bool
}

func (e *netError) Error() string {
	return e.msg
}

func (e *netError) Temporary() bool {
	return e.temporary
}

func (e *netError) Timeout() bool {
	return e.timeout
}

//CloseError represents a close message
type CloseError struct {
	//Code is defined in RFC 6455, section 11.7
	Code int
	//Text is the optional text payload
	Text string
}

func (e *CloseError) Error() string {
	s := []byte("ws: close")
	s = strconv.AppendInt(s, int64(e.Code), 10)
	switch e.Code {
	case CloseNormalClosure:
		s = append(s, " (normal)"...)
	case CloseGoingAway:
		s = append(s, " (going away)"...)
	case CloseProtocolError:
		s = append(s, " (protocol error)"...)
	case CloseUnsupportedData:
		s = append(s, " (unsupported data)"...)
	case CloseNoStatusReceived:
		s = append(s, " (no status)"...)
	case CloseAbnormalClosure:
		s = append(s, " (abnormal closure)"...)
	case CloseInvalidFramePayloadData:
		s = append(s, " (invalid payload data)"...)
	case ClosePolicyViolation:
		s = append(s, " (policy violation)"...)
	case CloseMessageTooBig:
		s = append(s, " (message too big)"...)
	case CloseMandatoryExtension:
		s = append(s, " (mandatory extension missing)"...)
	case CloseInternalServerErr:
		s = append(s, " (internal server error)"...)
	case CloseTLSHandshake:
		s = append(s, " (TLS hanshake error)"...)
	}
	if len(e.Text) != 0 {
		s = append(s, ": "...)
		s = append(s, e.Text...)
	}
	return string(s)
}

//IsCloseError returns boolean indicating whether the error is a *CloseError
//with one of the specified codes.
func IsCloseError(err error, codes ...int) bool {
	if e, ok := err.(*CloseError); ok {
		for _, code := range codes {
			if e.Code == code {
				return true
			}
		}
	}
	return false
}

//IsUnexpectedCloseError returns boolean indicating whether the error is a
//*CloseError with a code not in the list of expected codes.
func IsUnexpectedCloseError(err error, expectedCodes ...int) bool {
	if e, ok := err.(*CloseError); ok {
		for _, code := range expectedCodes {
			if e.Code == code {
				return false
			}
		}
	}
	return true
}

var (
	errWriteTimeout        = &netError{msg: "ws: write timeout", timeout: true, temporary: true}
	errUnexpectedEOF       = &CloseError{Code: CloseAbnormalClosure, Text: io.ErrUnexpectedEOF.Error()}
	errBadWriteOpCode      = errors.New("ws: bad write message type")
	errWriteClosed         = errors.New("ws: write closed")
	errInvalidControlFrame = errors.New("ws: invalid control frame")
)

func newMaskKey() [4]byte {
	n := rand.Uint32()
	return [4]byte{byte(n), byte(n >> 8), byte(n >> 16), byte(n >> 24)}
}

func hideTempErr(err error) error {
	if e, ok := err.(net.Error); ok && e.Temporary() {
		err = &netError{msg: e.Error(), timeout: e.Timeout()}
	}
	return err
}

func isControl(frameType int) bool {
	return frameType == CloseMessage ||
		frameType == PingMessage ||
		frameType == PongMessage
}

func isData(frameType int) bool {
	return frameType == TextMessage ||
		frameType == BinaryMessage
}

var validReceivedCloseCodes = map[int]bool{
	CloseNormalClosure:           true,
	CloseGoingAway:               true,
	CloseProtocolError:           true,
	CloseUnsupportedData:         true,
	CloseNoStatusReceived:        true,
	CloseAbnormalClosure:         false,
	CloseInvalidFramePayloadData: true,
	ClosePolicyViolation:         true,
	CloseMessageTooBig:           true,
	CloseMandatoryExtension:      true,
	CloseInternalServerErr:       true,
	CloseServiceRestart:          true,
	CloseTryAgainLater:           true,
	CloseTLSHandshake:            false,
}

func isValidReceivedCloseCode(code int) bool {
	return validReceivedCloseCodes[code] || (code >= 300 && code <= 4999)
}

//BufferPool represents a pool of buffers. The &sync.Pool type satisfes(满足) this
//interface. The type of the value stored in a pool is not specified.
type BufferPool interface {
	//Get gets a value from the pool or returns nil if the pool is empty.
	Get() interface{}
	//Put adds a value to the pool
	Put(interface{})
}
