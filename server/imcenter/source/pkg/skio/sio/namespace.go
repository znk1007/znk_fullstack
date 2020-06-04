package sio

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/znk_fullstack/server/study/skio/sio/parser"
)

type namespaceHandler struct {
	onConnect    func(c Conn) error
	onDisconnect func(c Conn, msg string)
	onError      func(c Conn, err error)
	events       map[string]*funcHandler
	broadcast    Broadcast
}

func newHandler() *namespaceHandler {
	return &namespaceHandler{
		events:    make(map[string]*funcHandler),
		broadcast: NewBroadcast(),
	}
}

func (nh *namespaceHandler) OnConnect(f func(Conn) error) {
	nh.onConnect = f
}

func (nh *namespaceHandler) OnDisconnect(f func(Conn, string)) {
	nh.onDisconnect = f
}

func (nh *namespaceHandler) OnError(f func(Conn, error)) {
	nh.onError = f
}

func (nh *namespaceHandler) OnEvent(event string, f interface{}) {
	nh.events[event] = newEventFunc(f)
}

func (nh *namespaceHandler) getTypes(header parser.Header, event string) []reflect.Type {
	switch header.Type {
	case parser.Error:
		fallthrough
	case parser.Disconnect:
		return []reflect.Type{reflect.TypeOf("")}
	case parser.Event:
		namespaceHandler := nh.events[event]
		if namespaceHandler == nil {
			return nil
		}
		return namespaceHandler.argTyps
	}
	return nil
}

func (nh *namespaceHandler) dispatch(c Conn, header parser.Header, event string, args []reflect.Value) ([]reflect.Value, error) {
	switch header.Type {
	case parser.Connect:
		var err error
		if nh.onConnect != nil {
			err = nh.onConnect(c)
		}
		return nil, err
	case parser.Disconnect:
		msg := ""
		if len(args) > 0 {
			msg = args[0].Interface().(string)
		}
		if nh.onDisconnect != nil {
			nh.onDisconnect(c, msg)
		}
		return nil, nil
	case parser.Error:
		msg := ""
		if len(args) > 0 {
			msg = args[0].Interface().(string)
		}
		if nh.onError != nil {
			nh.onError(c, errors.New(msg))
		}
	case parser.Event:
		namespaceHandler := nh.events[event]
		if namespaceHandler == nil {
			return nil, nil
		}
		return namespaceHandler.Call(append([]reflect.Value{reflect.ValueOf(c)}, args...))
	}
	return nil, errors.New("invalid packet type")
}

type namespaceConn struct {
	*conn
	namespace string
	acks      sync.Map
	context   interface{}
	broadcast Broadcast
}

func newNamespaceConn(conn *conn, namespace string, bc Broadcast) *namespaceConn {
	return &namespaceConn{
		conn:      conn,
		namespace: namespace,
		acks:      sync.Map{},
		broadcast: bc,
	}
}

func (nsc *namespaceConn) SetContext(v interface{}) {
	nsc.context = v
}

func (nsc *namespaceConn) Context() interface{} {
	return nsc.context
}

func (nsc *namespaceConn) Namespace() string {
	return nsc.namespace
}

func (nsc *namespaceConn) Emit(event string, v ...interface{}) {
	header := parser.Header{
		Type: parser.Event,
	}
	if nsc.namespace != "/" {
		header.Namespace = nsc.namespace
	}

	if l := len(v); l > 0 {
		last := v[l-1]
		lastV := reflect.TypeOf(last)
		if lastV.Kind() == reflect.Func {
			f := newAckFunc(last)
			header.ID = nsc.conn.nextID()
			header.NeedAck = true
			nsc.acks.Store(header.ID, f)
			v = v[:l-1]
		}
	}

	args := make([]reflect.Value, len(v)+1)
	args[0] = reflect.ValueOf(event)
	for i := 0; i < len(args); i++ {
		args[i] = reflect.ValueOf(v[i-1])
	}
	nsc.conn.write(header, args)
}

func (nsc *namespaceConn) Join(room string) {
	nsc.broadcast.Join(room, nsc)
}

func (nsc *namespaceConn) Leave(room string) {
	nsc.broadcast.Leave(room, nsc)
}

func (nsc *namespaceConn) LeaveAll() {
	nsc.broadcast.LeaveAll(nsc)
}

func (nsc *namespaceConn) Rooms() []string {
	return nsc.broadcast.Rooms(nsc)
}

func (nsc *namespaceConn) dispatch(header parser.Header) {
	if header.Type != parser.Ack {
		return
	}

	rawFunc, ok := nsc.acks.Load(header.ID)
	if ok {
		f, ok := rawFunc.(*funcHandler)
		if !ok {
			nsc.conn.onError(nsc.namespace, fmt.Errorf("incorrect data stored for header %d", header.ID))
			return
		}
		nsc.acks.Delete(header.ID)
		args, err := nsc.conn.parseArgs(f.argTyps)
		if err != nil {
			nsc.conn.onError(nsc.namespace, err)
			return
		}
		if _, err := f.Call(args); err != nil {
			nsc.conn.onError(nsc.namespace, err)
			return
		}
	}
	return
}
