package core

import (
	"testing"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

func TestHeader(t *testing.T) {
	tests := []struct {
		header pbs.Header
		bin    byte
	}{
		{
			pbs.Header{T: pbs.Type_connect},
			0,
		},
	}
	for _, test := range tests {
		t.Error("test bin byte == header type byte: ", test.bin == byte(test.header.T))
		t.Error("test type: ", test.header.T)
	}
}
