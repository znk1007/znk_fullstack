package core

import (
	"fmt"
	"io"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

type packetEncoder struct {
	dw dataWriter
}

// NewPacketEncoder 包编码器
func NewPacketEncoder(dw dataWriter) *packetEncoder {
	return &packetEncoder{
		dw: dw,
	}
}

func (pe *packetEncoder) NextWriter(dt pbs.DataType, pt pbs.PacketType) (io.WriteCloser, error) {
	w, err := pe.dw.NextWriter(dt, pt)
	if err != nil {
		return nil, err
	}
	var b [1]byte
	if dt == pbs.DataType_string {
		b[0] = packetTypeToASCIIByte(pt)
	} else {
		b[0] = packetTypeToBinaryByte(pt)
	}
	if _, err := w.Write(b[:]); err != nil {
		w.Close()
		return nil, err
	}
	return w, nil
}

type packetDecoder struct {
	dr dataReader
}

// NewPacketDecoder 包解码器
func NewPacketDecoder(dr dataReader) *packetDecoder {
	return &packetDecoder{
		dr: dr,
	}
}

func (pd *packetDecoder) NextReader() (pbs.DataType, pbs.PacketType, io.ReadCloser, error) {
	dt, _, r, err := pd.dr.NextReader()
	if err != nil {
		return 0, 0, nil, err
	}
	var b [1]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return 0, 0, nil, err
	}
	fmt.Println("packet coder b: ", b[0])
	return dt, byteToPacketType(b[0], dt), r, nil
}
