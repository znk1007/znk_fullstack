package base

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimestamp(t *testing.T) {
	should := assert.New(t)
	t1 := Timestamp()
	t2 := Timestamp()
	should.NotEmpty(t1)
	should.NotEmpty(t2)
	should.NotEqual(t1, t2)
}

func TestPacketType(t *testing.T) {
	at := assert.New(t)
	tests := []struct {
		b       byte
		ft      FrameType
		pt      PacketType
		strByte byte
		binByte byte
		str     string
	}{
		{0, FrameBinary, OPEN, '0', 0, "open"},
		{1, FrameBinary, CLOSE, '1', 1, "close"},
		{2, FrameBinary, PING, '2', 2, "ping"},
		{3, FrameBinary, PONG, '3', 3, "pong"},
		{4, FrameBinary, MESSAGE, '4', 4, "message"},
		{5, FrameBinary, UPGRADE, '5', 5, "upgrade"},
		{6, FrameBinary, NOOP, '6', 6, "noop"},

		{'0', FrameString, OPEN, '0', 0, "open"},
		{'1', FrameString, CLOSE, '1', 1, "close"},
		{'2', FrameString, PING, '2', 2, "ping"},
		{'3', FrameString, PONG, '3', 3, "pong"},
		{'4', FrameString, MESSAGE, '4', 4, "message"},
		{'5', FrameString, UPGRADE, '5', 5, "upgrade"},
		{'6', FrameString, NOOP, '6', 6, "noop"},
	}

	for _, test := range tests {
		bpt := ByteToPacketType(test.b, test.ft)
		at.Equal(test.pt, bpt)
		at.Equal(test.strByte, bpt.StringByte())
		at.Equal(test.binByte, bpt.BinaryByte(), "bpt: %v", bpt, " expect: %v", bpt.BinaryByte(), " want %v", test.binByte)
		at.Equal(test.str, bpt.String())
		at.Equal(test.str, PacketType(bpt).String())
	}
}

type fakeOpError struct {
	timeout   bool
	temporary bool
}

func (f fakeOpError) Error() string {
	return "fake error"
}

func (f fakeOpError) Timeout() bool {
	return f.timeout
}

func (f fakeOpError) Temporary() bool {
	return f.temporary
}

func TestOpError(t *testing.T) {
	at := assert.New(t)

	tests := []struct {
		url       string
		op        string
		err       error
		timeout   bool
		temporary bool
		errStr    string
	}{
		{"http://domain/abc", "post(write) to", io.EOF, false, false, "post(write) to http://domain/abc: EOF"},
		{"http://domain/abc", "get(read) from", io.EOF, false, false, "get(read) from http://domain/abc: EOF"},
		{"http://domain/abc", "post(write) to", fakeOpError{true, false}, true, false, "post(write) to http://domain/abc: fake error"},
		{"http://domain/abc", "get(read) from", fakeOpError{false, true}, false, true, "get(read) from http://domain/abc: fake error"},
	}
	for _, test := range tests {
		err := OpErr(test.url, test.op, test.err)
		e, ok := err.(*OpError)
		at.True(ok)
		at.Equal(test.timeout, e.Timeout())
		at.Equal(test.temporary, e.Temporary())
		at.Equal(test.errStr, e.Error())
	}
}
