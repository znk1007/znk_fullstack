package socket

import (
	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

// byteToDataType 比特转数据类型
func byteToDataType(b byte) pbs.DataType {
	return pbs.DataType(b)
}

// dataTypeToByte 数据类型转比特数据
func dataTypeToByte(dt pbs.DataType) byte {
	return byte(dt)
}

// packetTypeToASCIIByte 包类型转ascii比特
func packetTypeToASCIIByte(pt pbs.PacketType) byte {
	return byte(pt) + '0'
}

// packetTypeToBinaryByte 包类型转二进制比特
func packetTypeToBinaryByte(pt pbs.PacketType) byte {
	return byte(pt)
}

// byteToPacketType 比特转包类型
func byteToPacketType(b byte, dt pbs.DataType) pbs.PacketType {
	if dt == pbs.DataType_string {
		b -= '0'
	}
	return pbs.PacketType(b)
}
