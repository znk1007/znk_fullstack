package core

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

type testData struct {
	t  pbs.DataType
	bs []byte
}

type testPacket struct {
	dt pbs.DataType
	pt pbs.PacketType
	bs []byte
}

type testConnReader struct {
	data []testData
}

func newTestConnReader(data []testData) *testConnReader {
	return &testConnReader{
		data: data,
	}
}

func (tcr *testConnReader) NextReader() (pbs.DataType, pbs.PacketType, io.ReadCloser, error) {
	if len(tcr.data) == 0 {
		return pbs.DataType_string, 0, nil, io.EOF
	}
	d := tcr.data[0]
	tcr.data = tcr.data[1:]
	return d.t, 0, ioutil.NopCloser(bytes.NewReader(d.bs)), nil
}

type testConnWriter struct {
	data []testData
}

func newTestConnWriter() *testConnWriter {
	return &testConnWriter{}
}

func (tcw *testConnWriter) NextWriter(dt pbs.DataType, pt pbs.PacketType) (io.WriteCloser, error) {
	return newTestConnData(tcw, dt), nil
}

type testConnData struct {
	w *testConnWriter
	t pbs.DataType
	d *bytes.Buffer
}

func newTestConnData(tcw *testConnWriter, dt pbs.DataType) *testConnData {
	return &testConnData{
		w: tcw,
		t: dt,
		d: bytes.NewBuffer(nil),
	}
}

func (tcd *testConnData) Write(p []byte) (int, error) {
	return tcd.d.Write(p)
}

func (tcd *testConnData) Read(p []byte) (int, error) {
	return tcd.d.Read(p)
}

func (tcd *testConnData) Close() error {
	if tcd.w == nil {
		return nil
	}
	tcd.w.data = append(tcd.w.data, testData{
		t:  tcd.t,
		bs: tcd.d.Bytes(),
	})
	return nil
}

type testOneData struct {
	b byte
}

func (c *testOneData) Read(p []byte) (int, error) {
	p[0] = c.b
	return 1, nil
}

type testOneReader struct {
	dt  pbs.DataType
	tod *testOneData
}

func newOneReader() *testOneReader {
	return &testOneReader{
		dt: pbs.DataType_string,
		tod: &testOneData{
			b: packetTypeToASCIIByte(pbs.PacketType_message),
		},
	}
}

func (tor *testOneReader) NextReader() (pbs.DataType, pbs.PacketType, io.ReadCloser, error) {
	dt := tor.dt
	switch dt {
	case pbs.DataType_binary:
		tor.dt = pbs.DataType_string
		tor.tod.b = packetTypeToASCIIByte(pbs.PacketType_message)
	case pbs.DataType_string:
		tor.dt = pbs.DataType_binary
		tor.tod.b = packetTypeToBinaryByte(pbs.PacketType_message)
	}
	return dt, 0, ioutil.NopCloser(tor.tod), nil
}

type testOneDataDiscarder struct {
}

func (todd testOneDataDiscarder) Write(p []byte) (int, error) {
	return len(p), nil
}

func (todd testOneDataDiscarder) Close() error {
	return nil
}

type testDiscardWriter struct {
}

func (tdw *testDiscardWriter) NextWriter(dt pbs.DataType, pt pbs.PacketType) (io.WriteCloser, error) {
	return testOneDataDiscarder{}, nil
}

var tests = []struct {
	packets []testPacket
	data    []testData
}{
	{
		nil,
		nil,
	},
	{
		[]testPacket{
			{
				pbs.DataType_string,
				pbs.PacketType_open,
				[]byte{},
			},
		},
		[]testData{
			{
				pbs.DataType_string,
				[]byte("0"),
			},
		},
	},
	{
		[]testPacket{
			{
				pbs.DataType_string,
				pbs.PacketType_message,
				[]byte("hello 你好"),
			},
		},
		[]testData{
			{
				pbs.DataType_string,
				[]byte("4hello 你好"),
			},
		},
	},
	{
		[]testPacket{
			{
				pbs.DataType_binary,
				pbs.PacketType_message,
				[]byte("hello 你好"),
			},
		},
		[]testData{
			{
				pbs.DataType_binary,
				[]byte{0x04, 'h', 'e', 'l', 'l', 'o', ' ', 0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd},
			},
		},
	},
	{
		[]testPacket{
			{
				pbs.DataType_string,
				pbs.PacketType_open,
				[]byte{},
			},
			{
				pbs.DataType_binary,
				pbs.PacketType_message,
				[]byte("hello\n"),
			},
			{
				pbs.DataType_string,
				pbs.PacketType_message,
				[]byte("你好\n"),
			},
			{
				pbs.DataType_string,
				pbs.PacketType_ping,
				[]byte("probe"),
			},
		},
		[]testData{
			{
				pbs.DataType_string,
				[]byte("0"),
			},
			{
				pbs.DataType_binary,
				[]byte{0x04, 'h', 'e', 'l', 'l', 'o', '\n'},
			},
			{
				pbs.DataType_string,
				[]byte("4你好\n"),
			},
			{
				pbs.DataType_string,
				[]byte("2probe"),
			},
		},
	},
	{
		[]testPacket{
			{
				pbs.DataType_binary,
				pbs.PacketType_message,
				[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			},
			{
				pbs.DataType_string,
				pbs.PacketType_message,
				[]byte("hello"),
			},
			{
				pbs.DataType_string,
				pbs.PacketType_close,
				[]byte{},
			},
		},
		[]testData{
			{
				pbs.DataType_binary,
				[]byte{4, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			},
			{
				pbs.DataType_string,
				[]byte("4hello"),
			},
			{
				pbs.DataType_string,
				[]byte("1"),
			},
		},
	},
}

func TestPacketEncoder(t *testing.T) {
	for idx, test := range tests {
		tcw := newTestConnWriter()
		encoder := NewPacketEncoder(tcw)
		for _, p := range test.packets {
			nw, err := encoder.NextWriter(p.dt, p.pt)
			if err != nil {
				t.Error("encoder next writer err: ", err)
			}
			_, err = nw.Write(p.bs)
			if err != nil {
				t.Error("encoder write err: ", err)
			}
			err = nw.Close()
			if err != nil {
				t.Error("encoder close err: ", err)
			}
		}
		t.Errorf("idx: %d\n test data: %v\n enco data: %v", idx, test.data, tcw.data)
	}
}
