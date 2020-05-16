package ws

import (
	"net/http"
	"time"
)

//HandshakeError describes an error with the handshake from the peer
type HandshakeError struct {
	message string
}

func (e HandshakeError) Error() string { return e.message }

//Upgrader specifies paramters for upgrading an HTTP connection to a
//WebSocket connection.
type Upgrader struct {
	//HandshakeTimeout specifies the duration for the handshake to complete.
	HandshakeTimeout time.Duration

	//ReadBufferSize and WriteBufferSize specify I/O buffer sizes in bytes.
	//If a buffer size is zero, then buffers allocated by the HTTP server are used.
	//The I/O buffer sizes do not limit the size of the messages that can be sent
	// or received.
	ReadBufferSize, WriteBufferSize int

	//WriteBufferPool is pool of buffers for write operations. If the value
	//is not set, then write buffers are allocated to the connection for the
	//lifetime of the connection.
	//
	//A pool is most useful when the application has a modest volume of writes
	//across a large number of connections.
	//
	//Applications should use a single pool for each unique value of WriteBufferSize.
	WriteBufferPool BufferPool

	//Subprotocols specifies the server's supported protocols in order of
	//preference. If this field is not nil, then the Upgrade method negotiates
	//a subprotocol by selecting the first match in this list with a protocol
	//requested by the client. If there/s no match, then no protocol is
	//negotiated (the Sec-Websocket-Protocol header is not included in the
	//handshake response).
	Subprotocols []string

	//Error specifies the function for generating HTTP error responses. If Error
	//is nil, the http.Error is used to generate the HTTP response.
	Error func(w http.ResponseWriter, r *http.Request, status int, reason error)
}
