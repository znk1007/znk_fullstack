package primary

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	socket "github.com/znk_fullstack/golang/lib/utils/socket/protos/generated"
)

func TestConnParams(t *testing.T) {
	at := assert.New(t)
	param := socket.ConnParams{
		PingInterval: int64(time.Second * 10),
		PingTimeout:  int64(time.Second * 5),
		SID:          "vCcJKmYQcIf801WDAAAB",
		Upgrades:     []string{"websocket", "polling"},
	}
	out := param.String()
	tests := []struct {
		para socket.ConnParams
		out  string
	}{
		{
			param,
			out,
			// socket.ConnParams{
			// 	int64(time.Second * 10),
			// 	int64(time.Second * 5),
			// 	"vCcJKmYQcIf801WDAAAB",
			// 	[]string{"websocket", "polling"},
			// },
			// "{\"sid\":\"vCcJKmYQcIf801WDAAAB\",\"upgrades\":[\"websocket\",\"polling\"],\"pingInterval\":10000,\"pingTimeout\":5000}\n",
		},
	}
	for _, test := range tests {
		buf := bytes.NewBuffer(nil)

		_, err := WriteTo(test.para, buf) //test.para.WriteTo(buf)
		at.Nil(err)
		// at.Equal(test.out, buf.String())

		conn, err := ReadConnParams(buf)
		at.Nil(err)
		at.Equal(test.para, conn)
	}
}