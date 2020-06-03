package sio

import (
	"net/http"

	"github.com/znk_fullstack/server/study/skio/ngxio"
)

//Server for socket
type Server struct {
	handlers map[string]*namespaceHandler
	eio      *ngxio.Server
}

//NewServer create a socket server
func NewServer(opt *ngxio.Options) (*Server, error) {
	eio, err := ngxio.NewServer(opt)
	if err != nil {
		return nil, err
	}
	return &Server{
		handlers: make(map[string]*namespaceHandler),
		eio:      eio,
	}, nil
}

//Close closes server.
func (s *Server) Close() error {
	return s.eio.Close()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.eio.ServeHTTP(w, r)
}

//OnConnect set a handler function f
//to handle open event for namespace nsp
func (s *Server) OnConnect(nsp string, f func(Conn) error) {
	h := s.getNamespace(nsp, true)
	h.OnConnect(f)
}

//OnDisconnect set a handler function f to handle open event
//for namespace nsp.
func (s *Server) OnDisconnect(nsp string, f func(Conn, string)) {

}

func (s *Server) serveConn(c ngxio.Conn) {
	_, err := newConn(c, s.handlers)
	if err != nil {
		root := s.handlers[""]
		if root != nil && root.onError != nil {
			root.onError(nil, err)
		}
		return
	}
}

func (s *Server) getNamespace(nsp string, create bool) *namespaceHandler {
	if nsp == "/" {
		nsp = ""
	}
	ret, ok := s.handlers[nsp]
	if ok {
		return ret
	}
	if create {
		handler := newHandler()
		s.handlers[nsp] = handler
		return handler
	}
	return nil
}
