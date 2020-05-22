package packet

import (
	"testing"

	"github.com/znk_fullstack/server/study/skio/ngxio/base"
)

type Frame struct {
	ft   base.FrameType
	data []byte
}

type Packet struct {
	ft   base.FrameType
	pt   base.PacketType
	data []byte
}

var tests = []struct {
	packets []Packet
	frames  []Frame
}{
	{nil, nil},
	{
		[]Packet{
			{base.FrameString, base.OPEN, []byte{}},
		},
		[]Frame{
			{base.FrameString, []byte("0")},
		},
	},
	{
		[]Packet{
			{base.FrameString, base.MESSAGE, []byte("hello 你好")},
		},
		[]Frame{
			{base.FrameString, []byte("4hello 你好")},
		},
	},
	{
		[]Packet{
			{base.FrameBinary, base.MESSAGE, []byte("hello 你好")},
		},
		[]Frame{
			{base.FrameBinary, []byte{0x04, 'h', 'e', 'l', 'l', 'o', ' ', 0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd}},
		},
	},
	{
		[]Packet{
			{base.FrameString, base.OPEN, []byte{}},
			{base.FrameBinary, base.MESSAGE, []byte("hello\n")},
			{base.FrameString, base.MESSAGE, []byte("你好\n")},
			{base.FrameString, base.PING, []byte("probe")},
		},
		[]Frame{
			{base.FrameString, []byte("0")},
			{base.FrameBinary, []byte{0x04, 'h', 'e', 'l', 'l', 'o', '\n'}},
			{base.FrameString, []byte("4你好\n")},
			{base.FrameString, []byte("2probe")},
		},
	},
	{
		[]Packet{
			{base.FrameBinary, base.MESSAGE, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
			{base.FrameString, base.MESSAGE, []byte("hello")},
			{base.FrameString, base.CLOSE, []byte{}},
		},
		[]Frame{
			{base.FrameBinary, []byte{4, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
			{base.FrameString, []byte("4hello")},
			{base.FrameString, []byte("1")},
		},
	},
}

func TestEncoder(t *testing.T) {
	// at := assert.New(t)

	// for _, test := range tests {

	// }
}
