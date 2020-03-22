package userredis

import "testing"

func TestStart(t *testing.T) {
	readConf()
	err := Start()
	if err != nil {
		t.Error(err.Error())
	}
}
