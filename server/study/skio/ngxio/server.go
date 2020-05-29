package ngxio

import (
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/znk_fullstack/server/study/skio/bws"
	"github.com/znk_fullstack/server/study/skio/ngxio/base"
	"github.com/znk_fullstack/server/study/skio/ngxio/transport"
	"github.com/znk_fullstack/server/study/skio/ngxio/transport/polling"
	"github.com/znk_fullstack/server/study/skio/ngxio/transport/ws"
)

func defaultChecker(*http.Request) (http.Header, error) {
	return nil, nil
}

func defaultInitor(*http.Request, Conn) {}

//Options is options to create a server.
type Options struct {
	RequestChecker     func(*http.Request) (http.Header, error)
	ConnInitor         func(*http.Request, Conn)
	PingTimeout        time.Duration
	PingInterval       time.Duration
	Transports         []transport.Transport
	SessionIDGenerator SessionIDGenerator
}

func (o *Options) getRequestChecker() func(*http.Request) (http.Header, error) {
	if o != nil && o.RequestChecker != nil {
		return o.RequestChecker
	}
	return defaultChecker
}

func (o *Options) getConnInitor() func(*http.Request, Conn) {
	if o != nil && o.ConnInitor != nil {
		return o.ConnInitor
	}
	return defaultInitor
}

func (o *Options) getPingTimeout() time.Duration {
	if o != nil && o.PingTimeout != 0 {
		return o.PingInterval
	}
	return time.Minute
}

func (o *Options) getPingInterval() time.Duration {
	if o != nil && o.PingInterval != 0 {
		return o.PingInterval
	}
	return time.Second * 20
}

func (o *Options) getTransport() []transport.Transport {
	if o != nil && len(o.Transports) != 0 {
		return o.Transports
	}
	return []transport.Transport{
		polling.Default,
		ws.Default,
	}
}

func (o *Options) getSessionIDGenerator() SessionIDGenerator {
	if o != nil && o.SessionIDGenerator != nil {
		return o.SessionIDGenerator
	}
	return &defaultIDGenrator{}
}

//Server is server side
type Server struct {
	transports     *transport.Manager
	pingInterval   time.Duration
	pingTimeout    time.Duration
	sessions       *manager
	requestChecker func(*http.Request) (http.Header, error)
	connInitor     func(*http.Request, Conn)
	connChan       chan Conn
	closeOnce      sync.Once
}

//NewServer returns a server object.
func NewServer(ops *Options) (*Server, error) {
	t := transport.NewManager(ops.getTransport())
	return &Server{
		transports:     t,
		pingInterval:   ops.getPingInterval(),
		pingTimeout:    ops.getPingTimeout(),
		requestChecker: ops.getRequestChecker(),
		connInitor:     ops.getConnInitor(),
		sessions:       newManager(ops.getSessionIDGenerator()),
		connChan:       make(chan Conn, 1),
	}, nil
}

//Close close the server
func (s *Server) Close() error {
	s.closeOnce.Do(func() {
		close(s.connChan)
	})
	return nil
}

//Accept accepts a connection.
func (s *Server) Accept() (Conn, error) {
	c := <-s.connChan
	if c == nil {
		return nil, io.EOF
	}
	return c, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	sid := query.Get("sid")
	session := s.sessions.Get(sid)
	t := query.Get("transport")
	tspt := s.transports.Get(t)

	if tspt == nil {
		http.Error(w, "invalid tranport", http.StatusBadRequest)
		return
	}
	header, err := s.requestChecker(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	for k, v := range header {
		w.Header()[k] = v
	}
	if session == nil {
		if sid != "" {
			http.Error(w, "invalid sid", http.StatusBadRequest)
			return
		}
		conn, err := tspt.Accept(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		params := base.ConnParameters{
			PingInterval: s.pingInterval,
			PingTimeout:  s.pingTimeout,
			Upgrades:     s.transports.UpgradeFrom(t),
		}
		session, err = newSession(s.sessions, t, conn, params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		s.connInitor(r, session)

		go func() {
			w, err := session.nextWriter(base.FrameString, base.OPEN)
			if err != nil {
				session.Close()
				return
			}
			if _, err := session.params.WriteTo(w); err != nil {
				w.Close()
				session.Close()
				return
			}
			if err := w.Close(); err != nil {
				session.Close()
				return
			}
			s.connChan <- session
		}()
	}
	if session.Transport() != t {
		conn, err := tspt.Accept(w, r)
		if err != nil {
			//don't call http.Error() for HandshakeErrors because
			//they get handled by the websocket library internally.
			if _, ok := err.(bws.HandshakeError); !ok {
				http.Error(w, err.Error(), http.StatusBadGateway)
			}
			return
		}
		session.upgrade(t, conn)
		if handler, ok := conn.(http.Handler); ok {
			handler.ServeHTTP(w, r)
		}
		return
	}
	session.serveHTTP(w, r)
}
