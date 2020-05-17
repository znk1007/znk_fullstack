package ws

import (
	"io"
	"net/http"
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
