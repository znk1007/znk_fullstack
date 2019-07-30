package namespace

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"sync"

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

func (nh *namespaceHandler) dispatch(c Conn, header protos.Header, event string, args []reflect.Value) ([]reflect.Value, error) {
	switch header.Type {
	case protos.Header_connect:
		var err error
		if nh.onConnect != nil {
			err = nh.onConnect(c)
		}
		return nil, err
	case protos.Header_disconect:
		msg := ""
		if len(args) > 0 {
			msg = args[0].Interface().(string)
		}
		if nh.onDisconnect != nil {
			nh.onDisconnect(c, msg)
		}
		return nil, nil
	case protos.Header_error:
		msg := ""
		if len(args) > 0 {
			msg = args[0].Interface().(string)
		}
		if nh.onError != nil {
			nh.onError(errors.New(msg))
		}
	case protos.Header_event:
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
}

func newNamespaceConn(conn *conn, namespace string) *namespaceConn {
	return &namespaceConn{
		conn:      conn,
		namespace: namespace,
		acks:      sync.Map{},
	}
}

func (npc *namespaceConn) SetContext(v interface{}) {
	npc.context = v
}

func (npc *namespaceConn) Context() interface{} {
	return npc.context
}

func (npc *namespaceConn) Namespace() string {
	return npc.namespace
}

func (npc *namespaceConn) Emit(event string, v ...interface{}) {
	header := protos.Header{
		Type: protos.Header_event,
	}
	if npc.namespace != "/" {
		header.Namespace = npc.namespace
	}
	if l := len(v); l > 0 {
		last := v[l-1]
		lastV := reflect.TypeOf(last)
		if lastV.Kind() == reflect.Func {
			f := newAckFunc(lastV)
			header.ID = npc.conn.nextID()
			header.NeedAck = true
			npc.acks.Store(header.ID, f)
			v = v[:l-1]
		}
	}
	args := make([]reflect.Value, len(v)+1)
	args[0] = reflect.ValueOf(event)
	for idx := 0; idx < len(args); idx++ {
		args[idx] = reflect.ValueOf(v[idx-1])
	}
	npc.conn.write(header, args)
}

func (npc *namespaceConn) dispatch(header protos.Header) {
	if header.Type != protos.Header_ack {
		return
	}
	rawFunc, ok := npc.acks.Load(header.ID)
	if ok {
		f, ok := rawFunc.(*funcHandler)
		if !ok {
			npc.conn.onError(npc.namespace, fmt.Errorf(("incorrect data stored for header %c"), header.ID))
			return
		}
		npc.acks.Delete(header.ID)
		args, err := npc.conn.parseArgs(f.argTypes)
		if err != nil {
			npc.conn.onError(npc.namespace, err)
			return
		}
		if _, err := f.Call(args); err != nil {
			npc.conn.onError(npc.namespace, err)
			return
		}
	}
	return
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
	encoder    *core.Encoder
	decoder    *core.Decoder
	errorChan  chan errorMessage
	writeChan  chan writePacket
	quitChan   chan struct{}
	handlers   map[string]*namespaceHandler
	namespaces map[string]*namespaceConn
	closeOnce  sync.Once
	id         uint64
}

func newConn(c core.Conn, handlers map[string]*namespaceHandler) (*conn, error) {
	ret := &conn{
		Conn:       c,
		encoder:    core.NewEncoder(c),
		decoder:    core.NewDecoder(c),
		errorChan:  make(chan errorMessage),
		writeChan:  make(chan writePacket),
		quitChan:   make(chan struct{}),
		handlers:   handlers,
		namespaces: make(map[string]*namespaceConn),
	}
	if err := ret.connect(); err != nil {
		ret.Close()
		return nil, err
	}
	return ret, nil
}

func (c *conn) Close() error {
	var err error
	c.closeOnce.Do(func() {
		err = c.Conn.Close()
		close(c.quitChan)
	})
	return err
}

func (c *conn) connect() error {
	root := newNamespaceConn(c, "/")
	c.namespaces[""] = root
	header := protos.Header{
		Type: protos.Header_connect,
	}
	if err := c.encoder.Encode(header, nil); err != nil {
		return err
	}
	handler, ok := c.handlers[header.Namespace]
	go c.serveError()
	go c.serveWrite()
	go c.serveRead()
	if ok {
		handler.dispatch(root, header, "", nil)
	}
	return nil
}

func (c *conn) nextID() uint64 {
	c.id++
	return c.id
}

func (c *conn) write(header protos.Header, args []reflect.Value) {
	data := make([]interface{}, len(args))
	for i := range data {
		data[i] = args[i].Interface()
	}
	pkg := writePacket{
		header: header,
		data:   data,
	}
	select {
	case c.writeChan <- pkg:
	case <-c.quitChan:
		return
	}
}

func (c *conn) onError(namespace string, err error) {
	onErr := errorMessage{
		namespace: namespace,
		error:     err,
	}
	select {
	case c.errorChan <- onErr:
	case <-c.quitChan:
		return
	}
}

func (c *conn) parseArgs(types []reflect.Type) ([]reflect.Value, error) {
	return c.decoder.DecodeArgs(types)
}

func (c *conn) serveError() {
	defer c.Close()
	for {
		select {
		case <-c.quitChan:
			return
		case msg := <-c.errorChan:
			if handler := c.namespace(msg.namespace); handler != nil {
				if handler.onError != nil {
					handler.onError(msg.error)
				}
			}
		}
	}
}

func (c *conn) serveWrite() {
	defer c.Close()
	for {
		select {
		case <-c.quitChan:
			return
		case pkg := <-c.writeChan:
			if err := c.encoder.Encode(pkg.header, pkg.data); err != nil {
				c.onError(pkg.header.Namespace, err)
			}
		}
	}
}

func (c *conn) serveRead() {
	defer c.Close()
	var event string
	for {
		var header protos.Header
		if err := c.decoder.DecodeHeader(&header, &event); err != nil {
			c.onError("", err)
			return
		}
		if header.Namespace == "/" {
			header.Namespace = ""
		}
		switch header.Type {
		case protos.Header_ack:
			conn, ok := c.namespaces[header.Namespace]
			if !ok {
				c.decoder.DiscardLast()
				continue
			}
			conn.dispatch(header)
		case protos.Header_event:
			conn, ok := c.namespaces[header.Namespace]
			if !ok {
				c.decoder.DiscardLast()
				continue
			}
			handler, ok := c.handlers[header.Namespace]
			if !ok {
				c.decoder.DiscardLast()
				continue
			}
			types := handler.getTypes(header, event)
			args, err := c.decoder.DecodeArgs(types)
			if err != nil {
				c.onError(header.Namespace, err)
				return
			}
			ret, err := handler.dispatch(conn, header, event, args)
			if err != nil {
				c.onError(header.Namespace, err)
				return
			}
			if len(ret) > 0 {
				header.Type = protos.Header_ack
				c.write(header, ret)
			}
		case protos.Header_connect:
			if err := c.decoder.DiscardLast(); err != nil {
				c.onError(header.Namespace, err)
				return
			}
			conn, ok := c.namespaces[header.Namespace]
			if !ok {
				conn = newNamespaceConn(c, header.Namespace)
				c.namespaces[header.Namespace] = conn
			}
			handler, ok := c.handlers[header.Namespace]
			if ok {
				handler.dispatch(conn, header, "", nil)
			}
			c.write(header, nil)
		case protos.Header_disconect:
			types := []reflect.Type{
				reflect.TypeOf(""),
			}
			args, err := c.decoder.DecodeArgs(types)
			if err != nil {
				c.onError(header.Namespace, err)
				return
			}
			conn, ok := c.namespaces[header.Namespace]
			if !ok {
				c.decoder.DiscardLast()
				continue
			}
			delete(c.namespaces, header.Namespace)
			handler, ok := c.handlers[header.Namespace]
			if ok {
				handler.dispatch(conn, header, "", args)
			}
		}
	}
}

func (c *conn) namespace(npc string) *namespaceHandler {
	return c.handlers[npc]
}

// Server socket服务
type Server struct {
	handlers map[string]*namespaceHandler
	cs       *core.Server
}

// NewServer 创建服务
func NewServer(ops *core.Options) (*Server, error) {
	cs, err := core.NewServer(ops)
	if err != nil {
		return nil, err
	}
	return &Server{
		handlers: make(map[string]*namespaceHandler),
		cs:       cs,
	}, nil
}

// Close 关闭服务
func (s *Server) Close() error {
	return s.cs.Close()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.cs.ServeHTTP(w, r)
}

// OnConnect 连接
func (s *Server) OnConnect(ns string, f func(Conn) error) {
	h := s.getNamespace(ns)
	h.OnConnect(f)
}

// OnDisconnect 断开连接
func (s *Server) OnDisconnect(ns string, f func(Conn, string)) {
	h := s.getNamespace(ns)
	h.OnDisconnect(f)
}

// OnError 错误监听
func (s *Server) OnError(ns string, f func(error)) {
	h := s.getNamespace(ns)
	h.OnError(f)
}

// OnEvent 事件监听
func (s *Server) OnEvent(ns, event string, f interface{}) {
	h := s.getNamespace(ns)
	h.OnEvent(event, f)
}

// Serve 服务
func (s *Server) Serve() error {
	for {
		conn, err := s.cs.Accept()
		if err != nil {
			return err
		}
		go s.serveConn(conn)
	}
}

func (s *Server) serveConn(c core.Conn) {
	_, err := newConn(c, s.handlers)
	if err != nil {
		root := s.handlers[""]
		if root != nil && root.onError != nil {
			root.onError(err)
		}
		return
	}
}

func (s *Server) getNamespace(ns string) *namespaceHandler {
	if ns == "/" {
		ns = ""
	}
	ret, ok := s.handlers[ns]
	if ok {
		return ret
	}
	handler := newHandler()
	s.handlers[ns] = handler
	return handler
}
