package namespace

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"reflect"

	"github.com/znk_fullstack/golang/lib/utils/socket/core"

	protos "github.com/znk_fullstack/golang/lib/utils/socket/protos/generated"
)

type funcHandler struct {
	argTypes []reflect.Type
	f        reflect.Value
}

func newEventFunc(f interface{}) *funcHandler {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		panic("event handler must be a func")
	}
	ft := fv.Type()
	if ft.NumIn() < 1 || ft.In(0).Name() != "Conn" {
		panic("handler func should be like func(namespace.Conn, ...)")
	}
	argTypes := make([]reflect.Type, ft.NumIn()-1)
	for i := range argTypes {
		argTypes[i] = ft.In(i + 1)
	}
	if len(argTypes) == 0 {
		argTypes = nil
	}
	return &funcHandler{
		argTypes: argTypes,
		f:        fv,
	}
}

func newAckFunc(f interface{}) *funcHandler {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		panic("ack callback must be a func")
	}
	ft := fv.Type()
	argTypes := make([]reflect.Type, ft.NumIn())
	for i := range argTypes {
		argTypes[i] = ft.In(i)
	}
	if len(argTypes) == 0 {
		argTypes = nil
	}
	return &funcHandler{
		argTypes: argTypes,
		f:        fv,
	}
}

func (fh *funcHandler) Call(args []reflect.Value) (ret []reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("event call error: %s", r)
			}
		}
	}()
	ret = fh.f.Call(args)
	return
}

type namespaceHandler struct {
	onConnect    func(c Conn) error
	onDisconnect func(c Conn, msg string)
	onError      func(err error)
	events       map[string]*funcHandler
}

func newHandler() *namespaceHandler {
	return &namespaceHandler{
		events: make(map[string]*funcHandler),
	}
}

func (nh *namespaceHandler) OnConnect(f func(Conn) error) {
	nh.onConnect = f
}

func (nh *namespaceHandler) OnDisconnect(f func(Conn, string)) {
	nh.onDisconnect = f
}

func (nh *namespaceHandler) OnError(f func(error)) {
	nh.onError = f
}

func (nh *namespaceHandler) OnEvent(event string, f interface{}) {
	nh.events[event] = newEventFunc(f)
}

func (nh *namespaceHandler) getTypes(header protos.Header, event string) []reflect.Type {
	switch header.Type {
	case protos.Header_error:
		fallthrough
	case protos.Header_event:
		namespaceHandler := nh.events[event]
		if namespaceHandler == nil {
			return nil
		}
		return namespaceHandler.argTypes
	}
	return nil
}

// Conn 接口
type Conn interface {
	ID() string
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	Context() interface{}
	SetContext(v interface{})
	Namespace() string
	Emit(msg string, v ...interface{})
}

type errorMessage struct {
	namespace string
	error
}

type writePacket struct {
	header protos.Header
	data   []interface{}
}

type conn struct {
	core.Conn
	encoder   *core.Encoder
	decoder   *core.Decoder
	errorChan chan errorMessage
	writeChan chan writePacket
	quitChan  chan struct{}
	handlers  map[string]*namespaceHandler
}
