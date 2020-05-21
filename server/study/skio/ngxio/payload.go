package ngxio

import "io"

//byteReader interface for reading data
type byteReader interface {
	ReadByte() (byte, error)
	io.Reader
}

type readerFeeder interface {
	getReder() (io.Reader, bool, error)
	putReader(error) error
}

//decoder decode data sent from client
type decoder struct {
	feeder readerFeeder
	ft     frame
}
