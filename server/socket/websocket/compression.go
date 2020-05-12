package ws

import (
	"compress/flate"
	"io"
	"sync"
)

const (
	minCompressionLevel     = -2 //flate.HuffmanOnly not deined in go < 1.6
	maxCompressionLevel     = flate.BestCompression
	defaultCompressionLevel = 1
)

var (
	flateWriterPools [maxCompressionLevel-minCompressionLevel+1]sync.Pool
	flateReaderPool = sync.Pool{
		New: func() interface{} {
			return flate.NewReader(nil)
		},
	}
)

func compressNoContextTakeover(w io.WriteCloser, level int) io.WriteCloser {
	p := &
}