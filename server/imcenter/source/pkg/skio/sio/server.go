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
func (s *Server) OnConnect(namespace string, f func(Conn) error) {
	h := s.getNamespace(namespace, true)
	h.OnConnect(f)
}

//OnDisconnect set a handler function f to handle open event
//for namespace nsp.
func (s *Server) OnDisconnect(namespace string, f func(Conn, string)) {
	h := s.getNamespace(namespace, true)
	h.OnDisconnect(f)
}

//OnError set a handler function f to
//handle error for namespace nsp.
func (s *Server) OnError(namespace string, f func(Conn, error)) {
	h := s.getNamespace(namespace, true)
	h.OnError(f)
}

//OnEvent set a handler function f to
//handle event for namespace nsp.
func (s *Server) OnEvent(namespace, event string, f interface{}) {
	h := s.getNamespace(namespace, true)
	h.OnEvent(event, f)
}

//Serve serves socket server
func (s *Server) Serve() error {
	for {
		conn, err := s.eio.Accept()
		if err != nil {
			return err
		}
		go s.serveConn(conn)
	}
}

//JoinRoom joins given connection to the room
func (s *Server) JoinRoom(namespace, room string, connection Conn) bool {
	nspHandler := s.getNamespace(namespace, false)
	if nspHandler != nil {
		nspHandler.broadcast.Join(room, connection)
		return true
	}
	return false
}

//LeaveRoom leaves given connection from the room
func (s *Server) LeaveRoom(namespace, room string, connection Conn) bool {
	nspHandler := s.getNamespace(namespace, false)
	if nspHandler != nil {
		nspHandler.broadcast.Leave(room, connection)
		return true
	}
	return false
}

//LeaveAllRooms leaves given connection from all rooms
func (s *Server) LeaveAllRooms(namespace string, connection Conn) bool {
	nspHandler := s.getNamespace(namespace, false)
	if nspHandler != nil {
		nspHandler.broadcast.LeaveAll(connection)
		return true
	}
	return false
}

//ClearRoom clears the room
func (s *Server) ClearRoom(namespace string, room string) bool {
	nspHandler := s.getNamespace(namespace, false)
	if nspHandler != nil {
		nspHandler.broadcast.Clear(room)
		return true
	}
	return false
}

//BroadcastToRoom broadcasts given event & args to all the connections in the room
func (s *Server) BroadcastToRoom(namespace, room, event string, args ...interface{}) bool {
	nspHandler := s.getNamespace(namespace, false)
	if nspHandler != nil {
		nspHandler.broadcast.Send(room, event, args...)
		return true
	}
	return false
}

//RoomLen gives number of connections in the room
func (s *Server) RoomLen(namespace, room string) int {
	nspHandler := s.getNamespace(namespace, false)
	if nspHandler != nil {
		return nspHandler.broadcast.Len(room)
	}
	return -1
}

//Rooms gives list of all the rooms
func (s *Server) Rooms(namespace string) []string {
	nspHandler := s.getNamespace(namespace, false)
	if nspHandler != nil {
		return nspHandler.broadcast.Rooms(nil)
	}
	return nil
}

//ForEach map room connections
func (s *Server) ForEach(namespace, room string, f EachFunc) bool {
	nspHandler := s.getNamespace(namespace, false)
	if nspHandler != nil {
		nspHandler.broadcast.ForEach(room, f)
		return true
	}
	return false
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
