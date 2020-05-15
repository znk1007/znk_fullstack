package ws

import "io"

type broadcastMessage struct {
	payload  []byte
	prepared *PreparedMessage
}

type broadcastConn struct {
	conn  *Conn
	msgCh chan *broadcastMessage
}

//broadcastBench allows to run broadcast benchmarks.
//In every broadcast benchmark we create many connections, then send the same
//message into every connection and wait for all writes complete.
//This emulates an application where many connections listen to the same data.
//i.e. PUB/SUB scenarios with many subscribers in one channel.
type broadcastBench struct {
	w           io.Writer
	message     *broadcastMessage
	closeCh     chan struct{}
	doenCh      chan struct{}
	count       int32
	conns       []*broadcastConn
	compression bool
	usePrepared bool
}

func newBroadcastConn(c *Conn) *broadcastConn {
	return &broadcastConn{
		conn:  c,
		msgCh: make(chan *broadcastMessage),
	}
}
