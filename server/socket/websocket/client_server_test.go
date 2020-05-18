package ws

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

var cstUpgrader = Upgrader{
	Subprotocols:      []string{"p0", "p1"},
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		http.Error(w, reason.Error(), status)
	},
}

var cstDialer = Dialer{
	Subprotocols:     []string{"p1", "p2"},
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 30 * time.Second,
}

type cstHandler struct {
	*testing.T
}

func (t cstHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != cstPath {
		t.Logf("path=%v, want %v", r.URL.Path, cstPath)
		http.Error(w, "bad path", http.StatusBadRequest)
		return
	}
	if r.URL.RawQuery != cstRawQuery {
		t.Logf("query=%v, want %v", r.URL.RawQuery, cstRawQuery)
		http.Error(w, "bad path", http.StatusBadRequest)
		return
	}
	subprotos := Subprotocols(r)
	if !reflect.DeepEqual(subprotos, cstDialer.Subprotocols) {
		t.Logf("subprotos=%v, want %v", subprotos, cstDialer.Subprotocols)
		http.Error(w, "bad protocol", http.StatusBadRequest)
		return
	}
	ws, err := cstUpgrader.Upgrade(w, r, http.Header{"Set-Cookie": {"sessionID=1234"}})
	if err != nil {
		t.Logf("Upgrade: %v", err)
		return
	}
	defer ws.Close()

	if ws.SubProtocol() != "p1" {
		t.Logf("Subprotocol() = %s, want p1", ws.SubProtocol())
		ws.Close()
		return
	}
	op, rd, err := ws.NextReader()
	if err != nil {
		t.Logf("NextReader: %v", err)
		return
	}
	wr, err := ws.NextWriter(op)
	if err != nil {
		t.Logf("NextWriter: %v", err)
	}
	if _, err = io.Copy(wr, rd); err != nil {
		t.Logf("NextWriter: %v", err)
		return
	}
	if err := wr.Close(); err != nil {
		t.Logf("Close: %v", err)
		return
	}
}

type cstServer struct {
	*httptest.Server
	URL string
	t   *testing.T
}

const (
	cstPath       = "/a/b"
	cstRawQuery   = "x=y"
	cstRequestURI = cstPath + "?" + cstRawQuery
)

func newServer(t *testing.T) *cstServer {
	var s cstServer
	s.Server = httptest.NewServer(cstHandler{t})
	s.Server.URL += cstRequestURI
	s.URL = makeWsProto(s.Server.URL)
	return &s
}

func newTLSServer(t *testing.T) *cstServer {
	var s cstServer
	s.Server = httptest.NewTLSServer(cstHandler{t})
	s.Server.URL += cstRequestURI
	s.URL = makeWsProto(s.Server.URL)
	return &s
}

func makeWsProto(s string) string {
	return "ws" + strings.TrimPrefix(s, "http")
}

func sendRecv(t *testing.T, ws *Conn) {
	const message = "Hello World!"
	if err := ws.SetWriteDeadline(time.Now().Add(time.Second)); err != nil {
		t.Fatalf("SetWriteDeadline: %v", err)
	}
	if err := ws.WriteMessage(TextMessage, []byte(message)); err != nil {
		t.Fatalf("SetReadDeadline: %v", err)
	}
	_, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("ReadMessage: %v", err)
	}
	if string(p) != message {
		t.Fatalf("message=%s, want %s", p, message)
	}
}

func TestProxyDial(t *testing.T) {
	s := newServer(t)
	defer s.Close()

	surl, _ := url.Parse(s.Server.URL)

	cstDialer := cstDialer //make local copy for modification on next line.
	cstDialer.Proxy = http.ProxyURL(surl)

	connect := false
	origHandler := s.Server.Config.Handler

	//Capture the request Host header.
	s.Server.Config.Handler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "CONNECT" {
				connect = true
				w.WriteHeader(http.StatusOK)
				return
			}
			if !connect {
				t.Log("connect not received")
				http.Error(w, "connect not received", http.StatusMethodNotAllowed)
				return
			}
			origHandler.ServeHTTP(w, r)
		},
	)
	ws, _, err := cstDialer.Dial(s.URL, nil)
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer ws.Close()
	sendRecv(t, ws)
}

func TestProxyAuthorizationDial(t *testing.T) {
	s := newServer(t)
	defer s.Close()

	surl, _ := url.Parse(s.Server.URL)
	surl.User = url.UserPassword("username", "password")

	cstDialer := cstDialer // make local copy for modification on next line.
	cstDialer.Proxy = http.ProxyURL(surl)

	connect := false
	origHandler := s.Server.Config.Handler

	//Capture the request Host header.
	s.Server.Config.Handler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			proxyAuth := r.Header.Get("Proxy-Authorization")
			expectedProxyAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("username:password"))
			if r.Method == "CONNECT" && proxyAuth == expectedProxyAuth {
				connect = true
				w.WriteHeader(http.StatusOK)
				return
			}
			if !connect {
				t.Log("connect with proxy authorization not received")
				http.Error(w, "connect with proxy authorization not received", http.StatusMethodNotAllowed)
				return
			}
			origHandler.ServeHTTP(w, r)
		},
	)
	ws, _, err := cstDialer.Dial(s.URL, nil)
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer ws.Close()
	sendRecv(t, ws)
}

func TestDial(t *testing.T) {
	s := newServer(t)
	defer s.Close()

	ws, _, err := cstDialer.Dial(s.URL, nil)
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer ws.Close()
	sendRecv(t, ws)
}

func TestDialCookieJar(t *testing.T) {
	s := newServer(t)
	defer s.Close()

	jar, _ := cookiejar.New(nil)
	d := cstDialer
	d.Jar = jar

	u, _ := url.Parse(s.URL)

	switch u.Scheme {
	case "ws":
		u.Scheme = "http"
	case "wss":
		u.Scheme = "https"
	}

	cookies := []*http.Cookie{{Name: "gorilla", Value: "ws", Path: "/"}}
	d.Jar.SetCookies(u, cookies)

	ws, _, err := d.Dial(s.URL, nil)
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer ws.Close()

	var gorilla string
	var sessionID string
	for _, c := range d.Jar.Cookies(u) {
		if c.Name == "gorilla" {
			gorilla = c.Value
		}

		if c.Name == "sessionID" {
			sessionID = c.Value
		}
	}
	if gorilla != "ws" {
		t.Error("Cookie not present in jar.")
	}

	if sessionID != "1234" {
		t.Error("Set-Cookie not received from the server.")
	}

	sendRecv(t, ws)
}

func rootCAs(t *testing.T, s *httptest.Server) *x509.CertPool {
	certs := x509.NewCertPool()
	for _, c := range s.TLS.Certificates {
		roots, err := x509.ParseCertificates(c.Certificate[len(c.Certificate)-1])
		if err != nil {
			t.Fatalf("error parsing server's root cert: %v", err)
		}
		for _, root := range roots {
			certs.AddCert(root)
		}
	}
	return certs
}

func TestDialTLS(t *testing.T) {
	s := newTLSServer(t)
	defer s.Close()

	d := cstDialer
	d.TLSClientConfig = &tls.Config{RootCAs: rootCAs(t, s.Server)}
	ws, _, err := d.Dial(s.URL, nil)
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer ws.Close()
	sendRecv(t, ws)
}

func TestDialTimeout(t *testing.T) {
	s := newServer(t)
	defer s.Close()

	d := cstDialer
	d.HandshakeTimeout = -1
	ws, _, err := d.Dial(s.URL, nil)
	if err == nil {
		ws.Close()
		t.Fatalf("Dial: nil")
	}
}

//requireDeadlineNetConn fails the current test when Read or Write are called
//with no deadline.
type requireDeadlineNetConn struct {
	t                  *testing.T
	c                  net.Conn
	readDeadlineIsSet  bool
	writeDeadlineIsSet bool
}

func (c *requireDeadlineNetConn) SetDeadline(t time.Time) error {
	c.writeDeadlineIsSet = !t.Equal(time.Time{})
	c.readDeadlineIsSet = c.writeDeadlineIsSet
	return c.c.SetDeadline(t)
}

func (c *requireDeadlineNetConn) SetReadDeadline(t time.Time) error {
	c.writeDeadlineIsSet = !t.Equal(time.Time{})
	return c.c.SetDeadline(t)
}

func (c *requireDeadlineNetConn) SetWriteDeadline(t time.Time) error {
	c.writeDeadlineIsSet = !t.Equal(time.Time{})
	return c.c.SetDeadline(t)
}

func (c *requireDeadlineNetConn) Write(p []byte) (int, error) {
	if !c.writeDeadlineIsSet {
		c.t.Fatalf("write with no deadline")
	}
	return c.c.Write(p)
}

func (c *requireDeadlineNetConn) Read(p []byte) (int, error) {
	if !c.readDeadlineIsSet {
		c.t.Fatalf("read with no deadline")
	}
	return c.c.Read(p)
}

func (c *requireDeadlineNetConn) Close() error         { return c.c.Close() }
func (c *requireDeadlineNetConn) LocalAddr() net.Addr  { return c.c.LocalAddr() }
func (c *requireDeadlineNetConn) RemoteAddr() net.Addr { return c.c.RemoteAddr() }

func TestHandshakeTimeout(t *testing.T) {
	s := newServer(t)
	defer s.Close()

	d := cstDialer
	d.NetDial = func(n, a string) (net.Conn, error) {
		c, err := net.Dial(n, a)
		return &requireDeadlineNetConn{c: c, t: t}, err
	}
	ws, _, err := d.Dial(s.URL, nil)
	if err != nil {
		t.Fatal("Dial:", err)
	}
	ws.Close()
}

func TestHandshakeTimeoutInContext(t *testing.T) {
	s := newServer(t)
	defer s.Close()

	d := cstDialer
	d.HandshakeTimeout = 0
	d.NetDialConText = func(ctx context.Context, n, a string) (net.Conn, error) {
		netDialer := &net.Dialer{}
		c, err := netDialer.DialContext(ctx, n, a)
		return &requireDeadlineNetConn{c: c, t: t}, err
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	ws, _, err := d.DialContext(ctx, s.URL, nil)
	if err != nil {
		t.Fatal("Dial:", err)
	}
	ws.Close()
}

func TestDialBadScheme(t *testing.T) {
	s := newServer(t)
	defer s.Close()

	ws, _, err := cstDialer.Dial(s.Server.URL, nil)
	if err == nil {
		ws.Close()
		t.Fatalf("Dial: nil")
	}
}

func TestDialBadOrigin(t *testing.T) {
	s := newServer(t)
	defer s.Close()

	ws, resp, err := cstDialer.Dial(s.URL, http.Header{"Origin": {"bad"}})
	if err == nil {
		ws.Close()
		t.Fatalf("Dial: nil")
	}
	if resp == nil {
		t.Fatalf("resp=nil, err=%v", err)
	}
	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("status=%d, want %d", resp.StatusCode, http.StatusForbidden)
	}
}

func TestDialBadHeader(t *testing.T) {
	s := newServer(t)
	defer s.Close()

	for _, k := range []string{
		"Upgrade",
		"Connection",
		"Sec-Websocket-Key",
		"Sec-Websocket-Version",
		"Sec-Websocket-Protocol",
	} {
		h := http.Header{}
		h.Set(k, "bad")
		ws, _, err := cstDialer.Dial(s.URL, http.Header{"Origin": {"bad"}})
		if err == nil {
			ws.Close()
			t.Errorf("Dial with header %s returned nil", k)
		}
	}
}

func TestBadMethod(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws, err := cstUpgrader.Upgrade(w, r, nil)
		if err == nil {
			t.Errorf("handshake succeeded, expect fail")
			ws.Close()
		}
	}))
	defer s.Close()

	req, err := http.NewRequest("POST", s.URL, strings.NewReader(""))
	if err != nil {
		t.Fatalf("NewRequest returned error %v", err)
	}
	req.Header.Set("Connection", "upgrade")
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Sec-Websocket-Version", "13")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Do returned error %v", err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status = %d, want %d", resp.StatusCode, http.StatusMethodNotAllowed)
	}
}
