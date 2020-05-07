package skpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	skframe "github.com/znk_fullstack/server/socket/engine/tool/frame"
)

func TestPacketType(t *testing.T) {
	at := assert.New(t)
	tests := []struct {
		b   byte
		ft  skframe.FrameType
		pt  PacketType
		sb  byte
		bb  byte
		str string
	}{
		{0, skframe.FrameBinary, Open, '0', 0, "open"},
		{1, skframe.FrameBinary, Close, '1', 1, "close"},
		{2, skframe.FrameBinary, Ping, '2', 2, "ping"},
		{3, skframe.FrameBinary, Pong, '3', 3, "pong"},
		{4, skframe.FrameBinary, Message, '4', 4, "message"},
		{5, skframe.FrameBinary, Upgrade, '5', 5, "upgrade"},
		{6, skframe.FrameBinary, Noop, '6', 6, "noop"},

		{'0', skframe.FrameString, Open, '0', 0, "open"},
		{'1', skframe.FrameString, Close, '1', 1, "close"},
		{'2', skframe.FrameString, Ping, '2', 2, "ping"},
		{'3', skframe.FrameString, Pong, '3', 3, "pong"},
		{'4', skframe.FrameString, Message, '4', 4, "message"},
		{'5', skframe.FrameString, Upgrade, '5', 5, "upgrade"},
		{'6', skframe.FrameString, Noop, '6', 6, "noop"},
	}

	for _, test := range tests {
		btp := ByteToPacketType(test.b, test.ft)
		at.Equal(test.pt, btp)
		at.Equal(test.sb, btp.StringByte())
		at.Equal(test.bb, btp.BinaryByte())
		at.Equal(test.str, btp.String())
		at.Equal(test.str, PacketType(btp).String())
	}
}
