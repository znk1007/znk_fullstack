package payload

import "github.com/znk_fullstack/server/study/skio/ngxio/base"

type Packet struct {
	ft   base.FrameType
	pt   base.PacketType
	data []byte
}

var tests = []struct {
	supportBinary bool
	data          []byte
	packets       []Packet
}{
	{
		true,
		[]byte{0x00, 0x01, 0xff, '0'},
		[]Packet{
			{base.FrameString, base.OPEN, []byte{}},
		},
	},
}
