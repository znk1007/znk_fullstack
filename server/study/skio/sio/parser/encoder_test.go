package parser

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/znk_fullstack/server/study/skio/ngxio"
)

type fakeWriter struct {
	ft      ngxio.FrameType
	current *bytes.Buffer
	fts     []ngxio.FrameType
	bufs    []*bytes.Buffer
}

func (w *fakeWriter) NextWriter(ft ngxio.FrameType) (io.WriteCloser, error) {
	w.current = bytes.NewBuffer(nil)
	w.ft = ft
	return w, nil
}

func (w *fakeWriter) Write(p []byte) (int, error) {
	return w.current.Write(p)
}

func (w *fakeWriter) Close() error {
	w.fts = append(w.fts, w.ft)
	w.bufs = append(w.bufs, w.current)
	return nil
}

func TestEncoder(t *testing.T) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			should := assert.New(t)
			must := require.New(t)

			w := fakeWriter{}
			encoder := NewEncoder(&w)
			v := test.Var
			if test.Header.Type == Event {
				v = append([]interface{}{test.Event}, test.Var...)
			}
			err := encoder.Encode(test.Header, v)
			must.Nil(err)
			must.Equal(len(test.Datas), len(w.fts))
			must.Equal(len(test.Datas), len(w.bufs))
			for i := range w.fts {
				if i == 0 {
					should.Equal(ngxio.TEXT, w.fts[i])
					should.Equal(string(test.Datas[i]), string(w.bufs[i].Bytes()))
					continue
				}
				should.Equal(ngxio.BINARY, w.fts[i])
				should.Equal(test.Datas[i], w.bufs[i].Bytes())
			}
		})
	}
}

func TestAttachBuffer(t *testing.T) {
	tests := []struct {
		name    string
		data    interface{}
		max     uint64
		binarys [][]byte
	}{
		{"&Buffer", &Buffer{Data: []byte{1, 2}}, 1, [][]byte{{1, 2}}},
		{"[]interface{}{Buffer}", []interface{}{&Buffer{Data: []byte{1, 2}}}, 1, [][]byte{{1, 2}}},
		{"[]interface{}{Buffer,Buffer}", []interface{}{&Buffer{Data: []byte{1, 2}}}, 1, [][]byte{{1, 2}}},
		{"[]interface{}{Buffer,Buffer}", []interface{}{
			&Buffer{Data: []byte{1, 2}},
			&Buffer{Data: []byte{3, 4}},
		}, 2, [][]byte{{1, 2}, {3, 4}}},
		{"[1]interface{}{Buffer}", [...]interface{}{&Buffer{Data: []byte{1, 2}}}, 1, [][]byte{{1, 2}}},
		{"Struct{Buffer}", struct {
			Data *Buffer
			I    int
		}{
			&Buffer{Data: []byte{1, 2}},
			3,
		}, 1, [][]byte{{1, 2}}},
		{"map{Buffer}", map[string]interface{}{
			"data": &Buffer{Data: []byte{1, 2}},
			"i":    3,
		}, 1, [][]byte{{1, 2}}},
	}
	e := Encoder{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			should := assert.New(t)
			must := require.New(t)
			index := uint64(0)
			b, err := e.attachBuffer(reflect.ValueOf(test.data), &index)
			must.Nil(err)
			should.Equal(test.max, index)
			should.Equal(test.binarys, b)
		})
	}
}
