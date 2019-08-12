package core

import (
	"bytes"
	"html/template"
	"net"
	"net/http"
	"net/url"
	"time"
)

type pollingServer struct {
	*payload
	supportBinary bool
	remoteHeader  http.Header
	localAddr     addr
	remoteAddr    addr
	url           url.URL
	jsonp         string
}

func newPollingServer(r *http.Request) Conn {
	query := r.URL.Query()
	supportBinary := query.Get("base64") == ""
	jsonp := query.Get("j")
	if jsonp != "" {
		supportBinary = false
	}
	return &pollingServer{
		payload:       NewPayload(supportBinary),
		supportBinary: supportBinary,
		remoteHeader:  r.Header,
		localAddr:     addr{r.Host},
		remoteAddr:    addr{r.RemoteAddr},
		url:           *r.URL,
		jsonp:         jsonp,
	}
}

func (ps *pollingServer) URL() url.URL {
	return ps.url
}

func (ps *pollingServer) LocalAddr() net.Addr {
	return ps.localAddr
}

func (ps *pollingServer) RemoteAddr() net.Addr {
	return ps.remoteAddr
}

func (ps *pollingServer) RemoteHeader() http.Header {
	return ps.remoteHeader
}

func (ps *pollingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if jsonp := r.URL.Query().Get("j"); jsonp != "" {
			buf := bytes.NewBuffer(nil)
			if err := ps.payload.FlushOut(buf); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/javascript; charset=UTF-8")
			pl := template.JSEscapeString(buf.String())
			w.Write([]byte("___eio[" + jsonp + "](\""))
			w.Write([]byte(pl))
			w.Write([]byte("\");"))
			return
		}
		if ps.supportBinary {
			w.Header().Set("Content-Type", "application/octet-stream")
		} else {
			w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		}
		if err := ps.payload.FlushOut(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	case "POST":
		mime := r.Header.Get("Content-Type")
		supportBinary, err := mimeSupportBinary(mime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ps.payload.FeedIn(r.Body, supportBinary); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte("ok"))
		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}

func (ps *pollingServer) SetReadDeadline(t time.Time) error {
	return nil
}
func (ps *pollingServer) SetWriteDeadline(t time.Time) error {
	return nil
}
func (ps *pollingServer) ID() string {
	return ""
}
func (ps *pollingServer) Context() interface{} {
	return nil
}
func (ps *pollingServer) SetContext(v interface{}) {

}
func (ps *pollingServer) Namespace() string {
	return ""
}
func (ps *pollingServer) Emit(msg string, v ...interface{}) {

}
