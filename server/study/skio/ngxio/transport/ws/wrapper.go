package ws

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/znk_fullstack/server/study/skio/ngxio/base"
	websocket "github.com/znk_fullstack/server/study/skio/ws"
)

type wrapper struct {
	*websocket.Conn
	writeLock  *sync.Mutex
	readLocker *sync.Mutex
}

func newWrapper(conn *websocket.Conn) wrapper {
	return wrapper{
		Conn:       conn,
		writeLock:  new(sync.Mutex),
		readLocker: new(sync.Mutex),
	}
}

func (w wrapper) NextReader() (base.FrameType, io.ReadCloser, error) {
	w.readLocker.Lock()
	t, r, e := w.Conn.NextReader()
	//The wrapper remains locked until the returned ReadCloser is Closed.
	if e != nil {
		w.readLocker.Unlock()
		return 0, nil, e
	}
	switch t {
	case websocket.TextMessage:
		return base.FrameString, newRcWrapper(w.readLocker, r), nil
	case websocket.BinaryMessage:
		return base.FrameBinary, newRcWrapper(w.readLocker, r), nil
	}
	w.readLocker.Unlock()
	return 0, nil, transport.errin
}

type rcWrapper struct {
	nagTime *time.Timer
	quitNag chan struct{}
	l       *sync.Mutex
	io.Reader
}

func newRcWrapper(l *sync.Mutex, r io.Reader) rcWrapper {
	timer := time.NewTimer(30 * time.Second)
	q := make(chan struct{})
	go func() {
		select {
		case <-q:
		case <-timer.C:
			fmt.Println("Did you forget to Close() the ReadCloser from NextReader?")
		}
	}()
	return rcWrapper{
		nagTime: timer,
		quitNag: q,
		l:       l,
		Reader:  r,
	}
}
