package ws_test

import (
	"net/http"

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

}
