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

func TestControl(t *testing.T) {
	const message = "this is a ping/pong message"
	for _, isServer := range []bool{true, false} {
		for _, isWriteControl := range []bool{true, false} {
			name := fmt.Sprintf("s:%v, wc:%v", isServer, isWriteControl)
			var connBuf bytes.Buffer
			wc := newTestConn(nil, &connBuf, isServer)
			rc := newTestConn(&connBuf, nil, !isServer)
			if isWriteControl {
				wc.WriteControl(PongMessage, []byte(message), time.Now().Add(time.Second))
			} else {
				w, err := wc.NextWriter(PongMessage)
				if err != nil {
					t.Errorf("%s: wc.NextWriter() returned %v", name, err)
					continue
				}
				if _, err := w.Write([]byte(message)); err != nil {
					t.Errorf("%s: w.Write() returned %v", name, err)
					continue
				}
				if err := w.Close(); err != nil {
					t.Errorf("%s: w.Close() returned %v", name, err)
					continue
				}
				var actualMessage string
				rc.SetPongHandler(func(s string) error { actualMessage = s; return nil })
				rc.NextReader()
				if actualMessage != message {
					t.Errorf("%s: pong=%q, want %q", name, actualMessage, message)
					continue
				}
			}
		}
	}
}

type simpleBufferPool struct {
	v interface{}
}

func (p *simpleBufferPool) Get() interface{} {
	v := p.v
	p.v = nil
	return v
}

func (p *simpleBufferPool) Put(v interface{}) {
	p.v = v
}

func TestWriteBufferPool(t *testing.T) {
	const message = "Now is the time for all good people to come to the aid of the party."

	var buf bytes.Buffer
	var pool simpleBufferPool
	rc := newTestConn(&buf, nil, false)

	//Specify writeBufferSize smaller than message size to ensure that pooling
	//works with fragmented messages.
	wc := newConn(fakeNetConn{Writer: &buf}, true, 1024, len(message)-1, &pool, nil, nil)

	if wc.writeBuf != nil {
		t.Fatal("writeBuf not nil after create")
	}

	//Part 1: test NextWriter/Write/Close

	w, err := wc.NextWriter(TextMessage)
	if err != nil {
		t.Fatalf("wc.NextWriter() returned %v", err)
	}

	if wc.writeBuf == nil {
		t.Fatal("writeBuf is nil after NextWriter")
	}

	writeBufAddr := &wc.writeBuf[0]

	if _, err := io.WriteString(w, message); err != nil {
		t.Fatalf("io.WriteString(w, message) returned %v", err)
	}

	if err := w.Close(); err != nil {
		t.Fatalf("w.Close() returned %v", err)
	}

	if wc.writeBuf != nil {
		t.Fatal("writeBuf not nil after w.Close()")
	}

	if wpd, ok := pool.v.(writePoolData); !ok || len(wpd.buf) == 0 || &wpd.buf[0] != writeBufAddr {
		t.Fatal("writeBuf not returned to pool")
	}

	opCode, p, err := rc.ReadMessage()
	if opCode != TextMessage || err != nil {
		t.Fatalf("ReadMessage() returned %d, p, %v", opCode, err)
	}

	if s := string(p); s != message {
		t.Fatalf("message is %s, want %s", s, message)
	}

	//Part 2: Test WriteMessage.

	if err := wc.WriteMessage(TextMessage, []byte(message)); err != nil {
		t.Fatalf("wc.WriteMessage() returned %v", err)
	}

	if wc.writeBuf != nil {
		t.Fatal("writeBuf not nil after wc.WriteMessage()")
	}

	if wpd, ok := pool.v.(writePoolData); !ok || len(wpd.buf) == 0 || &wpd.buf[0] != writeBufAddr {
		t.Fatal("writeBuf not returned to pool after WriteMessage")
	}

	opCode, p, err = rc.ReadMessage()
	if opCode != TextMessage || err != nil {
		t.Fatalf("ReadMessage() returned %d, p, %v", opCode, err)
	}

	if s := string(p); s != message {
		t.Fatalf("message is %s, want %s", s, message)
	}
}
