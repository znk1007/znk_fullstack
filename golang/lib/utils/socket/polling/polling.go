package polling

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"mime"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"github.com/znk_fullstack/golang/lib/utils/socket/primary"

	"github.com/znk_fullstack/golang/lib/utils/socket/payload"

	socket "github.com/znk_fullstack/golang/lib/utils/socket/protos/generated"
)

func mimeSupportBinary(m string) (bool, error) {
	t, p, e := mime.ParseMediaType(m)
	if e != nil {
		return false, e
	}
	switch t {
	case "application/octet-stream":
		return true, nil
	case "text/plain":
		charset := strings.ToLower(p["charset"])
		if charset != "utf-8" {
			return false, errors.New("invalid charset")
		}
		return false, nil
	}
	return false, errors.New("invalid content-type")
}

// Addr 网络请求地址
type Addr struct {
	Host string
}

// Network 网络连接类型
func (a Addr) Network() string {
	return "tcp"
}

// String 端口
func (a Addr) String() string {
	return a.Host
}

// Transport 轮询传输
type Transport struct {
	Client *http.Client
}

// Default 默认传输
var Default = &Transport{
	Client: &http.Client{
		Timeout: time.Minute,
	},
}

// Name 传输名称
func (t *Transport) Name() string {
	return "polling"
}

// Accept 接收传输数据
func (t *Transport) Accept(w http.ResponseWriter, r *http.Request) (primary.Conn, error) {
	conn := newServerConn(r)
	return conn, nil
}

// Dial 拨号连接
func (t *Transport) Dial(u *url.URL, requestHeader http.Header) (primary.Conn, error) {
	query := u.Query()
	query.Set("transport", t.Name())
	u.RawQuery = query.Encode()
	return dial(t.Client, u, requestHeader)
}

type clientConn struct {
	*payload.Payload
	httpClient   *http.Client
	request      http.Request
	remoteHeader atomic.Value
}

func (c *clientConn) Open() (socket.ConnParams, error) {
	go c.getOpen()
	_, pt, r, err := c.NextReader()
	if err != nil {
		return socket.ConnParams{}, err
	}
	if pt != primary.Open {
		return socket.ConnParams{}, errors.New("invalid open")
	}
	conn, err := primary.ReadConnParams(r)
	if err != nil {
		return socket.ConnParams{}, err
	}
	err = r.Close()
	if err != nil {
		return socket.ConnParams{}, err
	}
	query := c.request.URL.Query()
	query.Set("sid", conn.SID)
	c.request.URL.RawQuery = query.Encode()
	go c.serverGet()
	go c.serverPost()
	return conn, nil
}

// dial 连接
func dial(client *http.Client, url *url.URL, requestHeader http.Header) (*clientConn, error) {
	if client == nil {
		client = &http.Client{}
	}
	req, err := http.NewRequest("", url.String(), nil)
	if err != nil {
		return nil, err
	}
	for k, v := range requestHeader {
		req.Header[k] = v
	}
	supportBinary := req.URL.Query().Get("base64") == ""
	if supportBinary {
		req.Header.Set("Content-Type", "application/octet-stream")
	} else {
		req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	}
	ret := &clientConn{
		Payload:    payload.New(supportBinary),
		httpClient: client,
		request:    *req,
	}
	return ret, nil
}

func (c *clientConn) URL() url.URL {
	return *c.request.URL
}

func (c *clientConn) LocalAddr() net.Addr {
	return Addr{""}
}

func (c *clientConn) RemoteAddr() net.Addr {
	return Addr{
		c.request.Host,
	}
}

func (c *clientConn) RemoteHeader() http.Header {
	ret := c.remoteHeader.Load()
	if ret == nil {
		return nil
	}
	return ret.(http.Header)
}

func (c *clientConn) Resume() {
	c.Payload.Resume()
	go c.serverGet()
	go c.serverPost()
}

// getOpen 打开
func (c *clientConn) getOpen() {
	req := c.request
	query := req.URL.Query()
	url := *req.URL
	req.URL = &url
	req.Method = "GET"
	query.Set("t", primary.NewSocketID().String())
	req.URL.RawQuery = query.Encode()
	resp, err := c.httpClient.Do(&req)
	if err != nil {
		c.Payload.Store("get", err)
		c.Close()
		return
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("invalid request: %s(%d)", resp.Status, resp.StatusCode)
	}
	var supportBinary bool
	if err == nil {
		m := resp.Header.Get("Content-Type")
		supportBinary, err = mimeSupportBinary(m)
	}
	if err != nil {
		c.Payload.Store("get", err)
		c.Close()
		return
	}
	c.remoteHeader.Store(resp.Header)
	if err = c.Payload.FeedIn(resp.Body, supportBinary); err != nil {
		return
	}
}

// serverGet GET请求
func (c *clientConn) serverGet() {
	req := c.request
	query := req.URL.Query()
	url := *req.URL
	req.URL = &url
	req.Method = "GET"
	for {
		query.Set("t", primary.NewSocketID().String())
		req.URL.RawQuery = query.Encode()
		resp, err := c.httpClient.Do(&req)
		if err != nil {
			c.Payload.Store("get", err)
			c.Close()
			return
		}
		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf("invalid request: %s(%d)", resp.Status, resp.StatusCode)
		}
		var supportBinary bool
		if err == nil {
			m := resp.Header.Get("Content-Type")
			supportBinary, err = mimeSupportBinary(m)
		}
		if err != nil {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
			c.Payload.Store("get", err)
			c.Close()
			return
		}
		if err = c.Payload.FeedIn(resp.Body, supportBinary); err != nil {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
			return
		}
		c.remoteHeader.Store(resp.Header)
	}
}

// serverPost POST请求
func (c *clientConn) serverPost() {
	var buf bytes.Buffer
	req := c.request
	url := *req.URL
	req.URL = &url
	query := url.Query()
	req.Method = "POST"
	req.Body = ioutil.NopCloser(&buf)
	for {
		buf.Reset()
		if err := c.Payload.FlushOut(&buf); err != nil {
			return
		}
		query.Set("t", primary.NewSocketID().String())
		req.URL.RawQuery = query.Encode()
		resp, err := c.httpClient.Do(&req)
		if err != nil {
			c.Payload.Store("post", err)
			c.Close()
			return
		}
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			c.Payload.Store("post", fmt.Errorf("invalid response %s(%d)", resp.Status, resp.StatusCode))
			c.Close()
			return
		}
		c.remoteHeader.Store(resp.Header)
	}
}

type serverConn struct {
	*payload.Payload
	supportBinary bool
	remoteHeader  http.Header
	localAddr     Addr
	remoteAddr    Addr
	url           url.URL
	jsonp         string
}

func newServerConn(r *http.Request) primary.Conn {
	query := r.URL.Query()
	supportBinary := query.Get("base64") == ""
	jsonp := query.Get("j")
	if jsonp != "" {
		supportBinary = false
	}
	return &serverConn{
		Payload:       payload.New(supportBinary),
		supportBinary: supportBinary,
		remoteHeader:  r.Header,
		localAddr:     Addr{r.Host},
		remoteAddr:    Addr{r.RemoteAddr},
		url:           *r.URL,
		jsonp:         jsonp,
	}
}

func (s *serverConn) URL() url.URL {
	return s.url
}

func (s *serverConn) LocalAddr() net.Addr {
	return s.localAddr
}

func (s *serverConn) RemoteAddr() net.Addr {
	return s.remoteAddr
}

func (s *serverConn) RemoteHeader() http.Header {
	return s.remoteHeader
}

func (s *serverConn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if jsonp := r.URL.Query().Get("j"); jsonp != "" {
			buf := bytes.NewBuffer(nil)
			if err := s.Payload.FlushOut(buf); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/javascript; charset=UTF-8")
			pl := template.JSEscapeString(buf.String())
			w.Write([]byte("__eio[" + jsonp + "](\""))
			w.Write([]byte(pl))
			w.Write([]byte("\");"))
			return
		}
		if s.supportBinary {
			w.Header().Set("Content-Type", "application/octet-stream")
		} else {
			w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		}
		if err := s.Payload.FlushOut(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	case "POST":
		m := r.Header.Get("Content-Type")
		supportBinary, err := mimeSupportBinary(m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := s.Payload.FeedIn(r.Body, supportBinary); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte("ok"))
		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}
