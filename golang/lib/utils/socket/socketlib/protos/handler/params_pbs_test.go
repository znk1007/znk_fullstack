package socket

import (
	"bytes"
	"testing"
	"time"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

func TestConnParameters(t *testing.T) {
	tests := []struct {
		param pbs.ConnParameters
		ret   string
	}{
		{
			pbs.ConnParameters{
				PingInterval: int64(time.Second * 10),
				PingTimeout:  int64(time.Second * 5),
				SID:          "abcdefjHiJKLmnopq09f",
				Upgrades:     []string{"websocket", "polling"},
			},
			"pingInterval:10000000000 pingTimeout:5000000000 sID:\"abcdefjHiJKLmnopq09f\" upgrades:\"websocket\" upgrades:\"polling\"",
		},
	}
	for _, test := range tests {
		buf := bytes.NewBuffer(nil)
		n, err := writeTo(test.param, buf)
		t.Log("len: ", n)
		if err != nil {
			t.Log("write err: ", err)
		}
		t.Log("write result: ", buf)
		t.Log("buf string: ", buf.String())
		params, err := readConnParams(buf)
		if err != nil {
			t.Error("read err: ", err)
		}
		t.Log("read result: ", params)
	}
}
