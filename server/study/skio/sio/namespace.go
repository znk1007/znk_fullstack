package sio

import (
	"errors"
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
