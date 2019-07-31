package socket

import (
	"io"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

// dataReader 数据读取器
type dataReader interface {
	NextReader() (pbs.DataType, io.ReadCloser, error)
}

// dataWriter 数据写入器
type dataWriter interface {
	NextWriter(dt pbs.DataType, pt pbs.PacketType) (io.WriteCloser, error)
}
