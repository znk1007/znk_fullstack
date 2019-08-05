package core

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"testing"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

type packet struct {
	dt   pbs.DataType
	pt   pbs.PacketType
	data []byte
}

var payloadTests = []struct {
	supportBinary bool
	data          []byte
	packets       []packet
}{
	{
		true,
		[]byte{0x00, 0x01, 0xff, '0'},
		[]packet{
			packet{
				pbs.DataType_string, 
				pbs.PacketType_open, 
				[]byte{}
			}
		},
	},
	{
		true,
		[]byte{0x00, 0x01, 0x03, 0xff, '4', 'h', 'e', 'l', 'l', 'o', ' ', 0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd},
		[]packet{
			packet{
				pbs.DataType_string, 
				pbs.PacketType_message, 
				[]byte("hello 你好")
			}
		},
	},
	{
		true,
		[]byte{0x00, 0x01, 0x03, 0xff, 0x04, 'h', 'e', 'l', 'l', 'o', ' ', 0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd},
		[]packet{
			packet{
				pbs.DataType_binary,
				pbs.PacketType_message,
				[]byte("hello 你好"),
			},
		},
	},
	{
		true,
		[]byte{
			0x01, 0x07, 0xff, 0x04, 'h', 'e', 'l', 'l', 'o', '\n',
			0x00, 0x08, 0xff, '4', 0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd, '\n',
			0x00, 0x06, 0xff, '2', 'p', 'r', 'o', 'b', 'e',
		},
		[]packet{
			packet{
				pbs.DataType_binary, 
				pbs.PacketType_message, 
				[]byte("hello\n")
			},
			packet{
				pbs.DataType_string, 
				pbs.PacketType_message, 
				[]byte("你好\n")
			},
			packet{pbs.DataType_string, 
				pbs.PacketType_ping, 
				[]byte("probe")
			},
		},
	},
	{
		false,
		[]byte("1:0"),
		[]packet{
			packet{
				pbs.DataType_string, 
				pbs.PacketType_open, 
				[]byte{}
			},
		},
	},
	{
		false,
		[]byte("13:4hello 你好"),
		[]packet{
			packet{
				pbs.DataType_string, 
				pbs.PacketType_message, 
				[]byte("hello 你好")
			},
		},
	},
	{
		false,
		[]byte("18:b4aGVsbG8g5L2g5aW9"),
		[]packet{
			packet{
				pbs.DataType_binary, 
				pbs.PacketType_message, 
				[]byte("hello 你好")
			},
		},
	},
	{
		false,
		[]byte("10:b4aGVsbG8K8:4你好\n6:2probe"),
		[]packet{
			packet{
				pbs.DataType_binary, 
				pbs.PacketType_message, 
				[]byte("hello\n")
			},
			packet{
				pbs.DataType_string, 
				pbs.PacketType_message, 
				[]byte("你好\n")
			},
			packet{
				pbs.DataType_string, 
				pbs.PacketType_ping, 
				[]byte("probe")
			},
		},
	},
}

type testReader struct {
	r io.Reader
}

func (tr testReader) Read(p []byte) (int, error) {
	return tr.r.Read(p)
}

type testReaderManager struct {
	data          []byte
	supportBinary bool
	returnError   error
	sendError     error
	getCounter    int
	addCounter    int
}

func (trm *testReaderManager) getReader() (io.Reader, bool, error) {
	trm.getCounter++
	return testReader{bytes.NewReader(trm.data)}, trm.supportBinary, trm.returnError
}

func (trm *testReaderManager) addReader(err error) error {
	trm.addCounter++
	trm.sendError = err
	return trm.returnError
}

func TestPayloadDecoder(t *testing.T) {
	for _, test := range payloadTests {
		manager := testReaderManager{
			data:          test.data,
			supportBinary: test.supportBinary,
		}
		d := payloadDecoder{
			readM: &manager,
		}
		var packets []packet
		for idx := 0; idx < len(test.packets); idx++ {
			dt, pt, fr, err := d.NextReader()
			if err != nil {
				t.Error("test payload decoder next reader err: ", err)
				break
			}
			data, err := ioutil.ReadAll(fr)
			if err != nil {
				t.Error("test payload decoder read all err: ", err)
			}
			pkt := packet{
				dt:   dt,
				pt:   pt,
				data: data,
			}
			err = fr.Close()
			if err != nil {
				t.Error("test payload decoder close err: ", err)
			}
			packets = append(packets, pkt)
		}
		t.Logf("\ntest packets:%v\nreal packets:%v", test.packets, packets)
		t.Logf("read manager get counter: %d", manager.getCounter)
		t.Errorf("read manager add counter: %d", manager.addCounter)
	}
}

func TestPayloadDecoderNextReaderError(t *testing.T) {
	manager := testReaderManager{
		data:          []byte{0x00, 0x01, 0xff, '0'},
		supportBinary: true,
	}
	d := payloadDecoder{
		readM: &manager,
	}
	targetErr := errors.New("error")
	manager.returnError = targetErr
	_, _, _, err := d.NextReader()
	t.Log("target error: ", targetErr)
	t.Error("next reader err: ", err)
}

type testWriterManager struct {
	w           io.Writer
	returnError error
	passinErr   error
}

func (twm *testWriterManager) getWriter() (io.Writer, error) {
	if oe, ok := twm.returnError.(Error); ok && oe.Temporary() {
		twm.returnError = nil
		return nil, oe
	}
	return twm.w, twm.returnError
}

func (twm *testWriterManager) addWriter(err error) error {
	twm.passinErr = err
	return twm.returnError
}

func TestPayloadEncoder(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	twm := &testWriterManager{
		w: buf,
	}
	for _, test := range payloadTests {
		buf.Reset()
		e := payloadEncoder{
			supportBinary: test.supportBinary,
			writeM:        twm,
		}
		for _, pkt := range test.packets {
			fw, err := e.NextWriter(pkt.dt, pkt.pt)
			if err != nil {
				t.Error("payload encoder next writer err: ", err)
			}
			t.Log("fw data: ", pkt.data)
			_, err = fw.Write(pkt.data)
			if err != nil {
				t.Error("payload encoder write err: ", err)
			}
			err = fw.Close()
			if err != nil {
				t.Error("payload encoder close err: ", err)
			}
		}
		t.Errorf("\ntest data :%v\nbufs bytes:%v", test.data, buf.Bytes())
	}

}
