package ws

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"testing"
	"testing/iotest"
	"time"
)

var _ net.Error = errWriteTimeout

type fakeAddr int

var (
	localAddr  = fakeAddr(1)
	remoteAddr = fakeAddr(2)
)

func (a fakeAddr) Network() string {
	return "net"
}

func (a fakeAddr) String() string {
	return "str"
}

type fakeNetConn struct {
	io.Reader
	io.Writer
}

func (c fakeNetConn) Close() error                       { return nil }
func (c fakeNetConn) LocalAddr() net.Addr                { return localAddr }
func (c fakeNetConn) RemoteAddr() net.Addr               { return remoteAddr }
func (c fakeNetConn) SetDeadline(t time.Time) error      { return nil }
func (c fakeNetConn) SetReadDeadline(t time.Time) error  { return nil }
func (c fakeNetConn) SetWriteDeadline(t time.Time) error { return nil }

func newTestConn(r io.Reader, w io.Writer, isServer bool) *Conn {
	return newConn(fakeNetConn{Reader: r, Writer: w}, isServer, 1024, 1024, nil, nil, nil)
}

func TestFraming(t *testing.T) {
	frameSizes := []int{
		0, 1, 2, 124, 125, 126, 127, 128, 129, 65534, 65535,
		//65536, 65537
	}
	var readChunkers = []struct {
		name string
		f    func(io.Reader) io.Reader
	}{
		{"half", iotest.HalfReader},
		{"one", iotest.OneByteReader},
		{"asis", func(r io.Reader) io.Reader { return r }},
	}
	writeBuf := make([]byte, 65537)
	for i := range writeBuf {
		writeBuf[i] = byte(i)
	}
	var writers = []struct {
		name string
		f    func(w io.Writer, n int) (int, error)
	}{
		{"iocopy", func(w io.Writer, n int) (int, error) {
			nn, err := io.Copy(w, bytes.NewReader(writeBuf[:n]))
			return int(nn), err
		}},
		{"write", func(w io.Writer, n int) (int, error) {
			return w.Write(writeBuf[:n])
		}},
		{"string", func(w io.Writer, n int) (int, error) {
			return io.WriteString(w, string(writeBuf[:n]))
		}},
	}

	for _, compress := range []bool{false, true} {
		for _, isServer := range []bool{true, false} {
			for _, chunker := range readChunkers {
				var connBuf bytes.Buffer
				wc := newTestConn(nil, &connBuf, isServer)
				rc := newTestConn(chunker.f(&connBuf), nil, !isServer)
				if compress {
					wc.newCompressionWriter = compressNoContextTakeover
					rc.newDecompressionReader = decompressNoContextTakeover
				}
				for _, n := range frameSizes {
					for _, writer := range writers {
						name := fmt.Sprintf("z:%v, s:%v, r:%s, n:%d w:%s", compress, isServer, chunker.name, n, writer.name)

						w, err := wc.NextWriter(TextMessage)
						if err != nil {
							t.Errorf("%s: wc.NextWriter() returned %v", name, err)
							continue
						}
						nn, err := writer.f(w, n)
						if err != nil || nn != n {
							t.Errorf("%s: w.Write(writeBuf[:n]) returned %d, %v", name, nn, err)
							continue
						}
						err = w.Close()
						if err != nil {
							t.Errorf("%s: w.Close() returned %v", name, err)
							continue
						}
						opCode, r, err := rc.NextReader()
						if err != nil || opCode != TextMessage {
							t.Errorf("%s: NextReader() returned %d, r, %v", name, opCode, err)
							continue
						}

						t.Logf("frame size: %d", n)
						rbuf, err := ioutil.ReadAll(r)
						if err != nil {
							t.Errorf("%s: ReadFull() returned rbuf, %v", name, err)
							continue
						}
						if len(rbuf) != n {
							t.Errorf("%s: len(rbuf) is %d, want %d", name, len(rbuf), n)
							continue
						}
						for i, b := range rbuf {
							if byte(i) != b {
								t.Errorf("%s: bad byte at offset %d", name, i)
							}
						}
					}
				}
			}
		}
	}
}
