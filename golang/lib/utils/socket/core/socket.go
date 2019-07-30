package core

import (
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/znk_fullstack/golang/lib/utils/socket/polling"
	"github.com/znk_fullstack/golang/lib/utils/socket/websocket"

	protos "github.com/znk_fullstack/golang/lib/utils/socket/protos/generated"

	"github.com/znk_fullstack/golang/lib/utils/socket/transport"

	"github.com/znk_fullstack/golang/lib/utils/socket/primary"
)

type writer struct {
	io.WriteCloser
	locker    *sync.RWMutex
	closeOnce sync.Once
}

func newWriter(w io.WriteCloser, locker *sync.RWMutex) *writer {
	return &writer{
		WriteCloser: w,
		locker:      locker,
	}
}

func (w *writer) Close() (err error) {
	w.closeOnce.Do(func() {
		w.locker.Lock()
		defer w.locker.Unlock()
		err = w.WriteCloser.Close()
	})
	return
}

func (w *writer) Write(bs []byte) (int, error) {
	w.locker.Lock()
	defer w.locker.Unlock()
	return w.WriteCloser.Write(bs)
}

type reader struct {
	io.ReadCloser
	locker    *sync.RWMutex
	closeOnce sync.Once
}

func newReader(r io.ReadCloser, locker *sync.RWMutex) *reader {
	return &reader{
		ReadCloser: r,
		locker:     locker,
	}
}

func (r *reader) Close() (err error) {
	r.closeOnce.Do(func() {
		r.locker.Lock()
		io.Copy(ioutil.Discard, r.ReadCloser)
		err = r.ReadCloser.Close()
		r.locker.Unlock()
	})
	return
}

// FrameType 数据类型
type FrameType primary.FrameType

const (
	// Text 文本类型数据
	Text = FrameType(primary.FrameString)
	// Binary 二进制类型数据
	Binary = FrameType(primary.FrameBinary)
)

// Conn 连接接口
type Conn interface {
	ID() string
	NextReader() (FrameType, io.ReadCloser, error)
	NextWriter(ft FrameType) (io.WriteCloser, error)
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	SetContext(v interface{})
	Context() interface{}
}

// Dialer 拨号器
type Dialer struct {
	Transports []transport.Transport
}

// Dial 拨号
func (d *Dialer) Dial(urlStr string, requestHeader http.Header) (Conn, error) {
	u, e := url.Parse(urlStr)
	if e != nil {
		return nil, e
	}
	query := u.Query()
	query.Set("EIO", "3")
	u.RawQuery = query.Encode()
	var conn primary.Conn
	for idx := len(d.Transports) - 1; idx >= 0; idx-- {
		if conn != nil {
			conn.Close()
		}
		t := d.Transports[idx]
		conn, e := t.Dial(u, requestHeader)
		if e != nil {
			continue
		}
		var params protos.ConnParams
		if p, ok := conn.(transport.Opener); ok {
			params, e = p.Open()
			if e != nil {
				continue
			}
		} else {
			var pt primary.PacketType
			var r io.ReadCloser
			_, pt, r, e = conn.NextReader()
			if e != nil {
				continue
			}
			func() {
				defer r.Close()
				if pt != primary.Open {
					e = errors.New("invalid open")
					return
				}
				params, e = primary.ReadConnParams(r)
				if e != nil {
					return
				}
			}()
		}
		if e != nil {
			continue
		}
		ret := &client{
			conn:      conn,
			params:    params,
			transport: t.Name(),
			close:     make(chan struct{}),
		}
		go ret.server()
		return ret, nil
	}
	return nil, e
}

type client struct {
	conn      primary.Conn
	params    protos.ConnParams
	transport string
	context   interface{}
	close     chan struct{}
	closeOnce sync.Once
}

func (c *client) SetContext(v interface{}) {
	c.context = v
}

func (c *client) Context() interface{} {
	return c.context
}

func (c *client) ID() string {
	return c.params.SID
}

func (c *client) Transport() string {
	return c.transport
}

func (c *client) Close() error {
	c.closeOnce.Do(func() {
		close(c.close)
	})
	return c.conn.Close()
}

func (c *client) NextReader() (FrameType, io.ReadCloser, error) {
	for {
		ft, pt, r, err := c.conn.NextReader()
		if err != nil {
			return 0, nil, err
		}
		switch pt {
		case primary.Pong:
			c.conn.SetReadDeadline(time.Now().Add(time.Duration(c.params.PingTimeout)))
		case primary.Close:
			c.Close()
			return 0, nil, io.EOF
		case primary.Message:
			return FrameType(ft), r, nil
		}
		r.Close()
	}
}

func (c *client) NextWriter(ft FrameType) (io.WriteCloser, error) {
	return c.conn.NextWriter(primary.FrameType(ft), primary.Message)
}

func (c *client) URL() url.URL {
	return c.conn.URL()
}

func (c *client) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *client) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *client) RemoteHeader() http.Header {
	return c.conn.RemoteHeader()
}

func (c *client) server() {
	defer c.conn.Close()
	for {
		select {
		case <-c.close:
			return
		case <-time.After(time.Duration(c.params.PingInterval)):
		}
		w, err := c.conn.NextWriter(primary.FrameString, primary.Ping)
		if err != nil {
			return
		}
		if err := w.Close(); err != nil {
			return
		}
		c.conn.SetWriteDeadline(time.Now().Add(time.Duration(c.params.PingTimeout)))
	}
}

// SessionIDGenerator 会话ID生成接口
type SessionIDGenerator interface {
	NewID() string
}

type defaultIDGenerator struct {
	nextID uint64
}

func (g *defaultIDGenerator) NewID() string {
	id := atomic.AddUint64(&g.nextID, 1)
	return strconv.FormatUint(id, 36)
}

type session struct {
	params        protos.ConnParams
	mamanger      *manager
	closeOnce     sync.Once
	context       interface{}
	upgradeLocker sync.RWMutex
	transport     string
	conn          primary.Conn
}

func newSession(m *manager, t string, conn primary.Conn, params protos.ConnParams) (*session, error) {
	params.SID = m.NewID()
	ret := &session{
		transport: t,
		conn:      conn,
		params:    params,
		mamanger:  m,
	}
	m.Add(ret)
	ret.setDeadline()
	return ret, nil
}

func (s *session) SetContext(v interface{}) {
	s.context = v
}

func (s *session) Context() interface{} {
	return s.context
}

func (s *session) ID() string {
	return s.params.SID
}

func (s *session) Transport() string {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()
	return s.transport
}

func (s *session) Close() error {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()
	s.closeOnce.Do(func() {
		s.mamanger.Remove(s.params.SID)
	})
	return s.conn.Close()
}

func (s *session) NextReader() (FrameType, io.ReadCloser, error) {
	for {
		ft, pt, r, e := s.nextReader()
		if e != nil {
			return 0, nil, e
		}
		switch pt {
		case primary.Ping:
			e := func() error {
				w, err := s.nextWriter(ft, primary.Pong)
				if err != nil {
					return err
				}
				io.Copy(w, r)
				return w.Close()
			}()
			r.Close()
			if e != nil {
				s.Close()
				return 0, nil, e
			}
			s.setDeadline()
		case primary.Close:
			r.Close()
			s.Close()
			return 0, nil, io.EOF
		case primary.Message:
			return FrameType(ft), r, nil
		}
		r.Close()
	}
}

func (s *session) NextWriter(ft FrameType) (io.WriteCloser, error) {
	return s.nextWriter(primary.FrameType(ft), primary.Message)
}

func (s *session) URL() url.URL {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()
	return s.conn.URL()
}

func (s *session) LocalAddr() net.Addr {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()
	return s.conn.LocalAddr()
}

func (s *session) RemoteAddr() net.Addr {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()
	return s.conn.RemoteAddr()
}

func (s *session) RemoteHeader() http.Header {
	s.upgradeLocker.RLock()
	defer s.upgradeLocker.RUnlock()
	return s.conn.RemoteHeader()
}

func (s *session) nextReader() (primary.FrameType, primary.PacketType, io.ReadCloser, error) {
	var ft primary.FrameType
	var pt primary.PacketType
	var r io.ReadCloser
	var err error
	for {
		s.upgradeLocker.RLock()
		ft, pt, r, err = s.conn.NextReader()
		if err != nil {
			s.upgradeLocker.RUnlock()
			if op, ok := err.(primary.Error); ok {
				if op.Temporary() {
					continue
				}
			}
			return 0, 0, nil, err
		}
		s.upgradeLocker.RUnlock()
		return ft, pt, newReader(r, &s.upgradeLocker), nil
	}
}

func (s *session) nextWriter(ft primary.FrameType, pt primary.PacketType) (io.WriteCloser, error) {
	for {
		s.upgradeLocker.RLock()
		w, err := s.conn.NextWriter(ft, pt)
		if err != nil {
			s.upgradeLocker.RUnlock()
			if op, ok := err.(primary.Error); ok {
				if op.Temporary() {
					continue
				}
			}
			return nil, err
		}
		s.upgradeLocker.RUnlock()
		return newWriter(w, &s.upgradeLocker), nil
	}
}

func (s *session) setDeadline() {
	deadline := time.Now().Add(time.Duration(s.params.PingTimeout))
	var conn primary.Conn
	for {
		s.upgradeLocker.RLock()
		same := conn == s.conn
		s.upgradeLocker.RUnlock()
		if same {
			return
		}
		s.conn.SetReadDeadline(deadline)
		s.conn.SetWriteDeadline(deadline)
	}
}

func (s *session) serverHTTP(w http.ResponseWriter, r *http.Request) {
	s.upgradeLocker.RLock()
	conn := s.conn
	s.upgradeLocker.RUnlock()
	if h, ok := conn.(http.Handler); ok {
		h.ServeHTTP(w, r)
	}
}

func (s *session) upgrade(transport string, conn primary.Conn) {
	go s.upgrading(transport, conn)
}

func (s *session) upgrading(t string, conn primary.Conn) {
	deadline := time.Now().Add(time.Duration(s.params.PingTimeout))
	conn.SetReadDeadline(deadline)
	conn.SetWriteDeadline(deadline)

	ft, pt, r, err := conn.NextReader()
	if err != nil {
		conn.Close()
		return
	}
	if pt != primary.Ping {
		conn.Close()
		return
	}
	if err := r.Close(); err != nil {
		conn.Close()
		return
	}
	w, err := conn.NextWriter(ft, pt)
	if err != nil {
		conn.Close()
		return
	}
	if _, err := io.Copy(w, r); err != nil {
		conn.Close()
		return
	}
	if err := w.Close(); err != nil {
		conn.Close()
		return
	}
	s.upgradeLocker.RLock()
	old := s.conn
	s.upgradeLocker.RUnlock()
	p, ok := old.(transport.Pauser)
	if !ok {
		conn.Close()
		return
	}
	p.Pasue()
	defer func() {
		if p != nil {
			p.Resume()
		}
	}()
	_, pt, r, err = conn.NextReader()
	if err != nil {
		conn.Close()
		return
	}
	if pt != primary.Upgrade {
		conn.Close()
		return
	}
	if err := r.Close(); err != nil {
		conn.Close()
		return
	}
	s.upgradeLocker.Lock()
	s.conn = conn
	s.transport = t
	s.upgradeLocker.Unlock()
	p = nil
	old.Close()
}

type manager struct {
	SessionIDGenerator
	s      map[string]*session
	locker sync.RWMutex
}

func newManager(g SessionIDGenerator) *manager {
	if g == nil {
		g = &defaultIDGenerator{}
	}
	return &manager{
		SessionIDGenerator: g,
		s:                  make(map[string]*session),
	}
}

func (m *manager) Add(s *session) {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.s[s.ID()] = s
}

func (m *manager) Get(sid string) *session {
	m.locker.Lock()
	defer m.locker.Unlock()
	return m.s[sid]
}

func (m *manager) Remove(sid string) {
	m.locker.Lock()
	defer m.locker.Unlock()
	if _, ok := m.s[sid]; !ok {
		return
	}
	delete(m.s, sid)
}

func defaultChecker(*http.Request) (http.Header, error) {
	return nil, nil
}

func defaultInitor(*http.Request, Conn) {}

// Options 连接选项
type Options struct {
	RequestChecker     func(*http.Request) (http.Header, error)
	ConnInitor         func(*http.Request, Conn)
	PingTimeout        time.Duration
	PingInterval       time.Duration
	Transports         []transport.Transport
	SessionIDGenerator SessionIDGenerator
}

// getRquestChecker 获取请求校验器
func (op *Options) getRquestChecker() func(*http.Request) (http.Header, error) {
	if op != nil && op.RequestChecker != nil {
		return op.RequestChecker
	}
	return defaultChecker
}

func (op *Options) getConnInitor() func(*http.Request, Conn) {
	if op != nil && op.ConnInitor != nil {
		return op.ConnInitor
	}
	return defaultInitor
}

func (op *Options) getPintTimeout() time.Duration {
	if op != nil && op.PingTimeout != 0 {
		return op.PingTimeout
	}
	return time.Minute
}

func (op *Options) getPingInterval() time.Duration {
	if op != nil && op.PingInterval != 0 {
		return op.PingInterval
	}
	return time.Second * 20
}

func (op *Options) getTransports() []transport.Transport {
	if op != nil && len(op.Transports) != 0 {
		return op.Transports
	}
	return []transport.Transport{
		polling.Default,
		websocket.Default,
	}
}

func (op *Options) getSessionIDGenerator() SessionIDGenerator {
	if op != nil && op.SessionIDGenerator != nil {
		return op.SessionIDGenerator
	}
	return &defaultIDGenerator{}
}

// Server 服务端对象
type Server struct {
	transports     *transport.Manager
	pingInterval   time.Duration
	pingTimeout    time.Duration
	sessions       *manager
	requestChecker func(*http.Request) (http.Header, error)
	connInitor     func(*http.Request, Conn)
	locker         sync.RWMutex
	connChan       chan Conn
	closeOnce      sync.Once
}

// NewServer 实例化服务端对象
func NewServer(ops *Options) (*Server, error) {
	tm := transport.NewManager(ops.getTransports())
	return &Server{
		transports:     tm,
		pingInterval:   ops.getPingInterval(),
		pingTimeout:    ops.getPintTimeout(),
		requestChecker: ops.getRquestChecker(),
		connInitor:     ops.getConnInitor(),
		sessions:       newManager(ops.getSessionIDGenerator()),
		connChan:       make(chan Conn, 1),
	}, nil
}

// Close 关闭服务
func (s *Server) Close() error {
	s.closeOnce.Do(func() {
		close(s.connChan)
	})
	return nil
}

// Accept 接收连接
func (s *Server) Accept() (Conn, error) {
	cc := <-s.connChan
	if cc == nil {
		return nil, io.EOF
	}
	return cc, nil
}

// ServeHTTP 连接服务端请求
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	sid := query.Get("sid")
	session := s.sessions.Get(sid)
	tpStr := query.Get("transport")
	tp := s.transports.Get(tpStr)
	if tp == nil {
		http.Error(w, "invalid transport", http.StatusBadRequest)
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
		conn, err := tp.Accept(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		params := protos.ConnParams{
			PingInterval: int64(s.pingInterval),
			PingTimeout:  int64(s.pingTimeout),
			Upgrades:     s.transports.UpgradeFrom(tpStr),
		}
		session, err = newSession(s.sessions, tpStr, conn, params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		s.connInitor(r, session)
		go func() {
			w, err := session.nextWriter(primary.FrameString, primary.Open)
			if err != nil {
				session.Close()
				return
			}
			if _, err := primary.WriteTo(session.params, w); err != nil {
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
	if session.Transport() != tpStr {
		conn, err := tp.Accept(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		session.upgrade(tpStr, conn)
		if handler, ok := conn.(http.Handler); ok {
			handler.ServeHTTP(w, r)
		}
		return
	}
	session.serverHTTP(w, r)
}
