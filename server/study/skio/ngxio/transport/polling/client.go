package polling

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sync/atomic"

	"github.com/znk_fullstack/server/study/skio/ngxio/base"
	"github.com/znk_fullstack/server/study/skio/ngxio/payload"
)

type clientConn struct {
	*payload.Payload

	httpClient   *http.Client
	request      http.Request
	remoteHeader atomic.Value
}

func dial(client *http.Client, url *url.URL, requestHeader http.Header) (*clientConn, error) {
	req, err := http.NewRequest("", url.String(), nil)
	if err != nil {
		return nil, err
	}
	if client == nil {
		client = &http.Client{}
	}
	for k, v := range requestHeader {
		req.Header[k] = v
	}
	b64 := req.URL.Query().Get("b64")
	supportBinary := len(b64) == 0
	if supportBinary {
		requestHeader.Set("Content-Type", "application/octet-stream")
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

func (cc *clientConn) Open() (base.ConnParameters, error) {
	go cc.getOpen()

	_, pt, r, err := cc.NextReader()
	if err != nil {
		return base.ConnParameters{}, err
	}
	if pt != base.OPEN {
		r.Close()
		return base.ConnParameters{}, errors.New("invalid open")
	}
	conn, err := base.ReadConnParameters(r)
	if err != nil {
		r.Close()
		return base.ConnParameters{}, err
	}
	err = r.Close()
	if err != nil {
		return base.ConnParameters{}, err
	}
	query := cc.request.URL.Query()
	query.Set("sid", conn.SID)
	cc.request.URL.RawQuery = query.Encode()

	go cc.serveGet()
	go cc.servePost()

	return conn, nil
}

func (cc *clientConn) URL() url.URL {
	return *cc.request.URL
}

func (cc *clientConn) LocalAddr() net.Addr {
	return Addr{""}
}

func (cc *clientConn) RemoteAddr() net.Addr {
	return Addr{cc.request.Host}
}

func (cc *clientConn) RemoteHeader() http.Header {
	ret := cc.remoteHeader.Load()
	if ret == nil {
		return nil
	}
	return ret.(http.Header)
}

func (cc *clientConn) Resume() {
	cc.Payload.Resume()
	go cc.serveGet()
	go cc.servePost()
}

func (cc *clientConn) servePost() {
	var buf bytes.Buffer
	req := cc.request
	url := *req.URL
	req.URL = &url
	query := url.Query()
	req.Method = "POST"
	req.Body = ioutil.NopCloser(&buf)
	for {
		buf.Reset()
		if err := cc.Payload.FlushOut(&buf); err != nil {
			return
		}
		query.Set("t", base.Timestamp())
		req.URL.RawQuery = query.Encode()
		resp, err := cc.httpClient.Do(&req)
		if err != nil {
			cc.Payload.Store("post", err)
			cc.Close()
			return
		}
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			cc.Payload.Store("post", fmt.Errorf("invalid response: %s(%d)", resp.Status, resp.StatusCode))
			cc.Close()
			return
		}
		cc.remoteHeader.Store(resp.Header)
	}
}

func (cc *clientConn) getOpen() {
	req := cc.request
	query := req.URL.Query()
	url := *req.URL
	req.URL = &url
	req.Method = "GET"
	query.Set("t", base.Timestamp())
	req.URL.RawQuery = query.Encode()
	resp, err := cc.httpClient.Do(&req)
	if err != nil {
		cc.Payload.Store("get", err)
		cc.Close()
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
		mime := resp.Header.Get("Content-Type")
		supportBinary, err = mimeSupportBinary(mime)
	}
	if err != nil {
		cc.Payload.Store("get", err)
		cc.Close()
		return
	}
	cc.remoteHeader.Store(resp.Header)
	if err = cc.Payload.FeedIn(resp.Body, supportBinary); err != nil {
		return
	}
}

func (cc *clientConn) serveGet() {
	req := cc.request
	query := req.URL.Query()
	url := *req.URL
	req.URL = &url
	req.Method = "GET"
	for {
		query.Set("t", base.Timestamp())
		req.URL.RawQuery = query.Encode()
		resp, err := cc.httpClient.Do(&req)
		if err != nil {
			cc.Payload.Store("get", err)
			cc.Close()
			return
		}
		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf("invalid request: %s(%d)", resp.Status, resp.StatusCode)
		}
		var supportBinary bool
		if err == nil {
			mime := resp.Header.Get("Content-Type")
			supportBinary, err = mimeSupportBinary(mime)
		}
		if err != nil {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
			cc.Payload.Store("get", err)
			cc.Close()
			return
		}
		if err = cc.Payload.FeedIn(resp.Body, supportBinary); err != nil {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
			return
		}
		cc.remoteHeader.Store(resp.Header)
	}
}
