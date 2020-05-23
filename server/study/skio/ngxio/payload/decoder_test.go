package payload

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeReader struct {
	r io.Reader
}

func (r fakeReader) Read(p []byte) (int, error) {
	return r.r.Read(p)
}

type fakeReaderFeeder struct {
	data          []byte
	supportBinary bool
	returnError   error
	sendError     error
	getCounter    int
	setCounter    int
}

func (f *fakeReaderFeeder) getReader() (io.Reader, bool, error) {
	f.getCounter++
	return fakeReader{bytes.NewReader(f.data)}, f.supportBinary, f.returnError
}

func (f *fakeReaderFeeder) setReader(err error) error {
	f.setCounter++
	f.sendError = err
	return f.returnError
}

func TestDecoder(t *testing.T) {
	at := assert.New(t)
	must := require.New(t)

	for _, test := range tests {

	}
}
