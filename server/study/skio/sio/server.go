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

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
