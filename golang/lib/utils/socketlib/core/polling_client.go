package core

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

	"github.com/znk_fullstack/golang/lib/utils/socketlib/protos/pbs"
)

type pollingClient struct {
	*payload
	httpClient   *http.Client
	request      http.Request
	remoteHeader atomic.Value
}

func dial(client *http.Client, url *url.URL, requestHeader http.Header) (*pollingClient, error) {
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
		req.Header.Set("Content-Type", "test/plain;charset=UTF-8")
	}
	ret := &pollingClient{
		payload:    NewPayload(supportBinary),
		request:    *req,
		httpClient: client,
	}
	return ret, nil
}

func (pc *pollingClient) Open() (pbs.ConnParameters, error) {
	go pc.getOpen()
	_, pt, r, err := pc.NextReader()
	if err != nil {
		return pbs.ConnParameters{}, err
	}
	if pt != pbs.PacketType_open {
		return pbs.ConnParameters{}, errors.New("invalid open")
	}
	conn, err := readConnParams(r)
	if err != nil {
		return pbs.ConnParameters{}, err
	}
	err = r.Close()
	if err != nil {
		return pbs.ConnParameters{}, err
	}
	query := pc.request.URL.Query()
	query.Set("sid", conn.SID)
	pc.request.URL.RawQuery = query.Encode()

	go pc.serveGet()
	go pc.servePost()
	return conn, nil
}

func (pc *pollingClient) URL() url.URL {
	return *pc.request.URL
}

func (pc *pollingClient) LocalAddr() net.Addr {
	return addr{""}
}

func (pc *pollingClient) RemoteAddr() net.Addr {
	return addr{pc.request.Host}
}

func (pc *pollingClient) RemoteHeader() http.Header {
	ret := pc.remoteHeader.Load()
	if ret == nil {
		return nil
	}
	return ret.(http.Header)
}

func (pc *pollingClient) Resume() {
	pc.payload.Resume()
	go pc.serveGet()
	go pc.servePost()
}

func (pc *pollingClient) servePost() {
	var buf bytes.Buffer
	req := pc.request
	url := *req.URL
	req.URL = &url
	query := url.Query()
	req.Method = "POST"
	req.Body = ioutil.NopCloser(&buf)
	for {
		buf.Reset()
		if err := pc.payload.FlushOut(&buf); err != nil {
			return
		}
		query.Set("t", NewSocketID().String())
		req.URL.RawQuery = query.Encode()
		resp, err := pc.httpClient.Do(&req)
		if err != nil {
			pc.payload.Store("post", err)
			pc.Close()
			return
		}
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			pc.payload.Store("post", fmt.Errorf("invalid response: %s(%d)", resp.Status, resp.StatusCode))
			pc.Close()
			return
		}
		pc.remoteHeader.Store(resp.Header)
	}
}

func (pc *pollingClient) getOpen() {
	req := pc.request
	query := req.URL.Query()
	url := *req.URL
	req.URL = &url
	req.Method = "GET"
	query.Set("t", NewSocketID().String())
	req.URL.RawQuery = query.Encode()
	resp, err := pc.httpClient.Do(&req)
	if err != nil {
		pc.payload.Store("get", err)
		pc.Close()
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
		pc.payload.Store("get", err)
		pc.Close()
		return
	}
	pc.remoteHeader.Store(resp.Header)
	if err = pc.payload.FeedIn(resp.Body, supportBinary); err != nil {
		return
	}
}

func (pc *pollingClient) serveGet() {
	req := pc.request
	query := req.URL.Query()
	url := *req.URL
	req.URL = &url
	req.Method = "GET"
	for {
		query.Set("t", NewSocketID().String())
		req.URL.RawQuery = query.Encode()
		resp, err := pc.httpClient.Do(&req)
		if err != nil {
			pc.payload.Store("get", err)
			pc.Close()
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
			pc.payload.Store("get", err)
			pc.Close()
			return
		}
		if err = pc.payload.FeedIn(resp.Body, supportBinary); err != nil {
			io.Copy(ioutil.Discard, resp.Body)
			return
		}
		pc.remoteHeader.Store(resp.Header)
	}
}
