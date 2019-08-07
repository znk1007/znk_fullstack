package core

import (
	"fmt"
	"testing"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

func TestPkgType(t *testing.T) {
	tests := []struct {
		dtb  byte
		dt   pbs.DataType
		ptb  byte
		pt   pbs.PacketType
		str  byte
		bin  byte
		desc string
	}{
		{1, pbs.DataType_binary, 0, pbs.PacketType_open, '0', 0, "open"},
		{1, pbs.DataType_binary, 1, pbs.PacketType_close, '1', 1, "open"},
		{1, pbs.DataType_binary, 2, pbs.PacketType_ping, '2', 2, "open"},
		{1, pbs.DataType_binary, 3, pbs.PacketType_pong, '3', 3, "open"},
		{1, pbs.DataType_binary, 4, pbs.PacketType_message, '4', 4, "open"},
		{1, pbs.DataType_binary, 5, pbs.PacketType_upgrade, '5', 5, "open"},
		{1, pbs.DataType_binary, 6, pbs.PacketType_noop, '6', 6, "open"},

		{0, pbs.DataType_string, '0', pbs.PacketType_open, '0', 0, "open"},
		{0, pbs.DataType_string, '1', pbs.PacketType_close, '1', 1, "close"},
		{0, pbs.DataType_string, '2', pbs.PacketType_ping, '2', 2, "ping"},
		{0, pbs.DataType_string, '3', pbs.PacketType_pong, '3', 3, "pong"},
		{0, pbs.DataType_string, '4', pbs.PacketType_message, '4', 4, "message"},
		{0, pbs.DataType_string, '5', pbs.PacketType_upgrade, '5', 5, "upgrade"},
		{0, pbs.DataType_string, '6', pbs.PacketType_noop, '6', 6, "noop"},
	}
	for _, test := range tests {
		dtb := dataTypeToByte(test.dt)
		// dt := byteToDataType(test.b)
		fmt.Println("test dtb: ", dtb)
		// fmt.Println("dtb == test dtb: ", dtb == test.dtb)
	}

	fmt.Println("------------------------")
	for _, test := range tests {
		bpt := byteToPacketType(test.ptb, test.dt)
		fmt.Println("test bpt: ", bpt)
		// bb := packetTypeToBinaryByte(bpt)
		// ab := packetTypeToASCIIByte(bpt)
		// pts := bpt.String()
		// fmt.Println("binary byte: ", bb)
		// fmt.Println("string byte: ", ab)
		// fmt.Println("bb == test bin: ", bb == test.bin)
		// fmt.Println("ab == test str: ", ab == test.str)
		// fmt.Println("pts == test desc: ", pts == test.desc)
	}
}
