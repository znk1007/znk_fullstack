package payload

import (
	"bytes"
	"io"
	"io/ioutil"
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
		feeder := fakeReaderFeeder{
			data:          test.data,
			supportBinary: test.supportBinary,
		}
		d := decoder{
			feeder: &feeder,
		}
		var packets []Packet

		for i := 0; i < len(test.packets); i++ {
			ft, pt, fr, err := d.NextReader()
			if err != nil {
				must.Equal(io.EOF, err)
				break
			}
			data, err := ioutil.ReadAll(fr)
			must.Nil(err)
			packet := Packet{
				ft:   ft,
				pt:   pt,
				data: data,
			}
			err = fr.Close()
			must.Nil(err)
			packets = append(packets, packet)
		}
		at.Equal(test.packets, packets)
		at.Equal(feeder.getCounter, 1)
		at.Equal(feeder.setCounter, 1)
	}
}
