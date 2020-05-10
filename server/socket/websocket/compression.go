package ws

import "compress/flate"

const (
	minCompressionLevel     = -2 //flate.HuffmanOnly not deined in go < 1.6
	maxCompressionLevel     = flate.BestCompression
	defaultCompressionLevel = 1
)