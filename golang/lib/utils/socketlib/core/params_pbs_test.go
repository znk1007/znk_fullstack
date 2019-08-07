package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/znk_fullstack/golang/lib/utils/socketlib/protos/pbs"
)

func TestConnParameters(t *testing.T) {
	tests := []struct {
		param pbs.ConnParameters
	}{
		{
			pbs.ConnParameters{
				PingInterval: int64(time.Second * 10),
				PingTimeout:  int64(time.Second * 5),
				SID:          "abcdefjHiJKLmnopq09f",
				Upgrades:     []string{"websocket", "polling"},
			},
		},
	}
	for _, test := range tests {

		// marshalByte, e := test.param.Marshal()
		// t.Log("marshal e: ", e)
		// t.Log("marshal byte: ", marshalByte)
		// var p pbs.ConnParameters
		// e = p.Unmarshal(marshalByte)
		// t.Log("unmarshal p: ", p)
		// t.Log("unmarshal e: ", e)

		buf := bytes.NewBuffer(nil)
		n, err := writeTo(test.param, buf)
		t.Log("len: ", n)
		if err != nil {
			t.Log("write err: ", err)
		}
		t.Log("write result: ", buf)
		params, err := readConnParams(buf)
		if err != nil {
			t.Error("read err: ", err)
		}
		t.Error("read result: ", params)
	}
}
