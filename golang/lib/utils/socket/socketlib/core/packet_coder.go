package core

import (
	"io"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

type encoder struct {
	w dataWriter
}

func newEncoder(w dataWriter) *encoder {
	return &encoder{
		w: w,
	}
}

func (e *encoder) NextWriter(dt pbs.DataType, pt pbs.PacketType) (io.WriteCloser, error) {
	w, err := e.w.NextWriter(dt, pt)
	if err != nil {
		return nil, err
	}
	var b [1]byte
	if dt == pbs.DataType_string {
		b[0] = packetTypeToASCIIByte(pt)
	} else {
		b[0] = packetTypeToBinaryByte(pt)
	}
	return w, nil
}
