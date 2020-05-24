package payload

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeWriterFeeder struct {
	w           io.Writer
	returnError error
	passinErr   error
}

func (f *fakeWriterFeeder) getWriter() (io.Writer, error) {
	if oe, ok := f.returnError.(Error); ok && oe.Temporary() {
		f.returnError = nil
		return nil, oe
	}
	return f.w, f.returnError
}

func (f *fakeWriterFeeder) setWriter(err error) error {
	f.passinErr = err
	return f.returnError
}

func TestEncoder(t *testing.T) {
	at := assert.New(t)
	must := require.New(t)
	buf := bytes.NewBuffer(nil)
	f := &fakeWriterFeeder{
		w: buf,
	}
	for _, test := range tests {
		buf.Reset()
		e := encoder{
			supportBinary: test.supportBinary,
			feeder:        f,
		}
		for _, packet := range test.packets {
			fw, err := e.NextWriter(packet.ft, packet.pt)
			must.Nil(err)
			_, err = fw.Write(packet.data)
			must.Nil(err)
			err = fw.Close()
			must.Nil(err)
		}
		at.Equal(test.data, buf.Bytes())
	}
}
