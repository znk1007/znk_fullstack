package ws_test

import (
	"log"
	"net/http"
	"testing"

	ws "github.com/znk_fullstack/server/socket/websocket"
)

var (
	c   *ws.Conn
	req *http.Request
)

//	The websocket.IsUnexpectedCloseError function is useful for identifying
//	application and protocol errors.
//
//	This server application works with a client application running in the
//	browser. The client application does not explicitly close the websocket.
//	The only expected close message from the client has the code
//	ws.CloseGoingAway. All other close messages are likely the result of an
//	application or protocol error and are logged to aid debugging.
func ExampleIsUnexpectedCloseError() {
	for {
		messageType, p, err := c.ReadMessage()
		if err != nil {
			if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway) {
				log.Printf("error: %v, user-agent: %v", err, req.Header.Get("User-Agent"))
			}
			return
		}
		processMessage(messageType, p)
	}
}

func processMessage(mt int, p []byte) {}

// 	TestX prevents godoc from showing this entire file in the example.
//	Remove this function when a second example is added.
func TestX(t *testing.T) {}
