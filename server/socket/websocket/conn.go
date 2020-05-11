package ws

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"strconv"
	"sync"
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

//writePoolData is the type added to the write buffer pool.
//This wrapper is used to prevent applications from peeking at and depending on the values
//added to the pool
type writePoolData struct {
	buf []byte
}

//Conn represents a ws connection.
type Conn struct {
	conn        net.Conn
	isServer    bool
	subprotocol string

	//Write fields
	mu            chan struct{} //used as mutex to protect write to conn
	writeBuf      []byte        //frame is constructed in this buffer.
	writePool     BufferPool
	writeBufSize  int
	writeDeadline time.Time
	writer        io.WriteCloser //the current writer returned to the application
	isWriting     bool           //for best-effort concurrent write detection

	writeErrMu sync.Mutex
	writeErr   error

	enableWriteCompression bool
	compressionLevel       int
	compressionWriter      func(io.WriteCloser, int) io.WriteCloser

	//Read fields
	reader  io.ReadCloser //the current reader returned to the application
	readErr error
	br      *bufio.Reader
	// bytes remaining in current frame.
	// set setReadRemaining to safely update this value and prevent overflow
	readRemaining int64
	readFinal     bool  //true the current message has more frames.
	readLength    int64 //Message size.
	readLimit     int64 //Maximum message size.
	readMaskKey   [4]byte
	handlePong    func(string) error
	handlePing    func(string) error
	handleClose   func(int, string) error
	readErrCount  int
	messageReader *messageReader //the current low-level reader

	readDecompress         bool //whether last read frame had RSV1 set
	newDecompressionReader func(io.Reader) io.ReadCloser
}

func newConn(conn net.Conn, isServer bool, readBufferSize, writeBufferSize int, writeBufferPool BufferPool, br *bufio.Reader, writeBuf []byte) *Conn {
	if br == nil {
		if readBufferSize == 0 {
			readBufferSize = defaultReadBufferSize
		} else if readBufferSize < maxControlFramePayloadSize {
			//must be large enough for control frame
			readBufferSize = maxControlFramePayloadSize
		}
		br = bufio.NewReaderSize(conn, readBufferSize)
	}
	if writeBufferSize <= 0 {
		writeBufferSize = defaultWriteBufferSize
	}
	writeBufferSize += maxFrameHeaderSize
	if writeBuf == nil && writeBufferPool == nil {
		writeBuf = make([]byte, writeBufferSize)
	}
	mu := make(chan struct{}, 1)
	mu <- struct{}{}
	c := &Conn{
		isServer:               isServer,
		br:                     br,
		conn:                   conn,
		mu:                     mu,
		readFinal:              true,
		writeBuf:               writeBuf,
		writePool:              writeBufferPool,
		writeBufSize:           writeBufferSize,
		enableWriteCompression: true,
		compressionLevel:       defaultCompressionLevel,
	}
	c.SetCloseHandler(nil)
	c.SetPingHandler(nil)
	c.SetPongHandler(nil)
	return c
}

//setReadRemaining
func (c *Conn) setReadRemaining(n int64) error {
	if n < 0 {
		return ErrReadLimit
	}
	c.readRemaining = n
	return nil
}

//SubProtocol returns the negotiated protocol for the connection.
func (c *Conn) SubProtocol() string {
	return c.subprotocol
}

//Close closes the underlying network connection without sending or waiting
//for a close message.
func (c *Conn) Close() error {
	return c.conn.Close()
}

//LocalAddr returns the local network address.
func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

//RemoteAddr returns the remote network address.
func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

//writeFatal write fatal when error
func (c *Conn) writeFatal(err error) error {
	err = hideTempErr(err)
	c.writeErrMu.Lock()
	if c.writeErr == nil {
		c.writeErr = err
	}
	c.writeErrMu.Unlock()
	return err
}

//read read data and returns result data
func (c *Conn) read(n int) ([]byte, error) {
	p, err := c.br.Peek(n)
	if err == io.EOF {
		err = errUnexpectedEOF
	}
	c.br.Discard(len(p))
	return p, err
}

func (c *Conn) write(frameType int, deadline time.Time, buf0, buf1 []byte) error {

}

//WriteControl writes a control message with the given deadline.
//The allowed message types are CloseMessage, PingMessage and PongMessage.
func (c *Conn) WriteControl(messageType int, data []byte, deadline time.Time) error {
	if !isControl(messageType) {
		return errBadWriteOpCode
	}
	if len(data) > maxControlFramePayloadSize {
		return errInvalidControlFrame
	}
	b0 := byte(messageType) | finalBit
	b1 := byte(len(data))
	if !c.isServer {
		b1 |= maskBit
	}
	buf := make([]byte, 0, maxFrameHeaderSize+maxControlFramePayloadSize)
	buf = append(buf, b0, b1)
	if c.isServer {
		buf = append(buf, data...)
	} else {
		key := newMaskKey()
		buf = append(buf, key[:]...)
		buf = append(buf, data...)
		maskBytes(key, 0, buf[6:])
	}
	d := 1000 * time.Hour
	if !deadline.IsZero() {
		d = deadline.Sub(time.Now())
		if d < 0 {
			return errWriteTimeout
		}
	}

	timer := time.NewTimer(d)
	select {
	case <-c.mu:
		timer.Stop()
	case <-timer.C:
		return errWriteTimeout
	}
	defer func() {
		c.mu <- struct{}{}
	}()
	c.writeErrMu.Lock()
	err := c.writeErr
	c.writeErrMu.Unlock()
	if err != nil {
		return err
	}
	c.conn.SetWriteDeadline(deadline)
	_, err = c.conn.Write(buf)
	if err != nil {
		return c.writeFatal(err)
	}
	if messageType == CloseMessage {
		c.writeFatal(ErrCloseSent)
	}
	return err
}

//CloseHandler returns the current close handler
func (c *Conn) CloseHandler() func(code int, text string) error {
	return c.handleClose
}

//SetCloseHandler sets the handler for close messages received from the peer.
//The code argument to h is the received close code or CloseNoStatusReceived
//if the close message is empty. The default close handler sends a close
//message back to the peer.
//
//The handler function is called from the NextReader, ReadMessage and message
//reader Read methods. The application must read the connection to process
//close messages as described in the section on Control Message above.
//
//The connection read methods return a CloseError when a close message is received.
//Most applications should handle close messages as part of their
//normal error handling. Applications should only set a close handler when the
//application must perform some action before sending a close message back to
//the peer.
func (c *Conn) SetCloseHandler(h func(code int, text string) error) {
	if h == nil {
		h = func(code int, text string) error {
			msg := FormatCloseMessage(code, "")
			c.WriteControl(CloseMessage, msg, time.Now().Add(writeWait))
			return nil
		}
	}
	c.handleClose = h
}

//PingHandler returns the current ping handler
func (c *Conn) PingHandler() func(appData string) error {
	return c.handlePing
}

//SetPingHandler sets the handler for ping messages received from the peer.
//The appData argument to h is the PING message application data.
//The default ping handler sends a pong to the peer.
//
//The handler function is called from the NextReader, ReadMessage and message
//reader Read methods.
//The application must read the connection to process ping messages as described
//in the section on Control Messages above.
func (c *Conn) SetPingHandler(h func(appData string) error) {
	if h == nil {
		h = func(msg string) error {
			err := c.WriteControl(PongMessage, []byte(msg), time.Now().Add(writeWait))
			if err == ErrCloseSent {
				return nil
			} else if e, ok := err.(net.Error); ok && e.Temporary() {
				return nil
			}
			return err
		}
	}
	c.handlePing = h
}

//PongHandler returns the current pong handler
func (c *Conn) PongHandler() func(appData string) error {
	return c.handlePong
}

//SetPongHandler sets the handler for pong messages received from the peer.
//The appData argument to h is the PONG message application data.
//The default pong handler does nothing.
//
//The handler function is called from the NextReader, ReadMessage and message
//reader Read methods. The application must read the connection to process
//pong messages as described in the section on Control Messages above.
func (c *Conn) SetPongHandler(h func(appData string) error) {
	if h == nil {
		h = func(string) error {
			return nil
		}
	}
	c.handlePong = h
}

//FormatCloseMessage formats closeCode and text as a ws close message.
//An empty message is returned for code CloseNorStatusReceicved.
func FormatCloseMessage(closeCode int, text string) []byte {
	if closeCode == CloseNoStatusReceived {
		//Return empty message because it's illegal to send CloseNoStatusReceived.
		//Return non-nil value in case application checks for nil.
		return []byte{}
	}
	buf := make([]byte, 2+len(text))
	binary.BigEndian.PutUint16(buf, uint16(closeCode))
	copy(buf[2:], text)
	return buf
}

type messageReader struct {
	c *Conn
}
