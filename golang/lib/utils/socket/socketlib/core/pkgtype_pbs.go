package core

import (
	"fmt"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

// byteToDataType 比特转数据类型
func byteToDataType(b byte) pbs.DataType {
	fmt.Println("byteToDataType b: ", pbs.DataType(b))
	return pbs.DataType(b)
}

// dataTypeToByte 数据类型转比特数据
func dataTypeToByte(dt pbs.DataType) byte {
	fmt.Println("dt: ", dt)
	fmt.Println("byte dt:", byte(dt))
	return byte(dt)
}

// packetTypeToASCIIByte 包类型转ascii比特
func packetTypeToASCIIByte(pt pbs.PacketType) byte {
	fmt.Println("packetTypeToBinaryByte: ", byte(pt)+'0')
	return byte(pt) + '0'
}

// packetTypeToBinaryByte 包类型转二进制比特
func packetTypeToBinaryByte(pt pbs.PacketType) byte {
	fmt.Println("packetTypeToBinaryByte: ", byte(pt))
	return byte(pt)
}

// byteToPacketType 比特转包类型
func byteToPacketType(b byte, dt pbs.DataType) pbs.PacketType {
	if dt == pbs.DataType_string {
		b -= '0'
		fmt.Println("byteToPacketType string: ", b)
	} else {
		fmt.Println("byteToPacketType binary: ", b)
	}

	return pbs.PacketType(b)
}
