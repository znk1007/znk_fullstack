package polling

import (
	"bytes"
	"net"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/znk_fullstack/server/study/skio/ngxio/base"
	"github.com/znk_fullstack/server/study/skio/ngxio/payload"
)

type serverConn struct {
	*payload.Payload
	transport     *Transport
	supportBinary bool

	remoteHeader http.Header
	localAddr    Addr
	remoteAddr   Addr
	url          url.URL
	jsonp        string
}

func newServerConn(t *Transport, r *http.Request) base.Conn {
	query := r.URL.Query()
	supportBinary := query.Get("b64") == ""
	jsonp := query.Get("j")
	if len(jsonp) != 0 {
		supportBinary = false
	}
	return &serverConn{
		Payload:       payload.New(supportBinary),
		transport:     t,
		supportBinary: supportBinary,
		remoteHeader:  r.Header,
		localAddr:     Addr{r.Host},
		remoteAddr:    Addr{r.RemoteAddr},
		url:           *r.URL,
		jsonp:         jsonp,
	}
}

func (sc *serverConn) URL() url.URL {
	return sc.url
}

func (sc *serverConn) LocalAddr() net.Addr {
	return sc.localAddr
}

func (sc *serverConn) RemoteAddr() net.Addr {
	return sc.remoteAddr
}

func (sc *serverConn) RemoteHeader() http.Header {
	return sc.remoteHeader
}

func (sc *serverConn) SetHeaders(w http.ResponseWriter, r *http.Request) {
	userAgent := r.UserAgent()
	if strings.Contains(userAgent, ";MSIE") || strings.Contains(userAgent, "Trident/") {
		w.Header().Set("X-XSS-Protection", "0")
	}

	//jst in case the default behaviour gets changed and it has to handle an origin check
	checkOrigin := Default.CheckOrigin
	if sc.transport.CheckOrigin != nil {
		checkOrigin = sc.transport.CheckOrigin
	}

	if checkOrigin != nil && checkOrigin(r) {
		j := r.URL.Query().Get("j")
		isPolling := len(j) == 0
		if isPolling {
			origin := r.Header.Get("Origin")
			if len(origin) == 0 {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			} else {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		}
	}
}

func (sc *serverConn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	j := r.URL.Query().Get("j")
	switch r.Method {
	case "OPTIONS":
		if len(j) == 0 {
			sc.SetHeaders(w, r)
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(200)
		}
	case "GET":
		sc.SetHeaders(w, r)
		if jsonp := j; len(jsonp) != 0 {
			buf := bytes.NewBuffer(nil)
			if err := sc.Payload.FlushOut(buf); err != nil {
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
		if sc.supportBinary {
			w.Header().Set("Content-Type", "application/octet-stream")
		} else {
			w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		}
		if err := sc.Payload.FlushOut(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	case "POST":
		sc.SetHeaders(w, r)
		mime := r.Header.Get("Content-Type")
		supportBinary, err := mimeSupportBinary(mime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := sc.Payload.FeedIn(r.Body, supportBinary); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte("ok"))
		return
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}
