package ngxio

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/znk_fullstack/server/study/skio/ngxio/transport"
	"github.com/znk_fullstack/server/study/skio/ngxio/transport/polling"
)

func TestEnginePolling(t *testing.T)  {
	should := assert.New(t)
	must := require.New(t)

	srv,err := NewServer(nil)
	must.Nil(err)
	defer srv.Close()
	httpSrv := httptest.NewServer(srv)
	defer httpSrv.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		should := assert.New(t)
		must := require.New(t)

		conn, err := srv.Accept()
		must.Nil(err)
		defer conn.Close()

		ft,r,err := conn.NextReader()
		must.Nil(err)
		should.Equal(TEXT, ft)
		b,err:=ioutil.ReadAll(r)
		must.Nil(err)
		should.Equal("hello你好", string(b))
		err = r.Close()
		must.Nil(err)

		w,err := conn.NextWriter(BINARY)
		must.Nil(err)
		_,err = w.Write([]byte{1,2,3,4})
		must.Nil(err)
		err = w.Close()
		must.Nil(err)
	}()

	dialer := Dialer{
		Transports: []transport.Transport{polling.Default},
	}
	header := http.Header{}
	header.Set("X-EIO-Test","client")

	cnt,err := dialer.Dial(httpSrv.URL, header)
	must.Nil(err)

	w,err := cnt.NextWriter(TEXT)
	must.Nil(err)
	_,err = w.Write([]byte("hello你好"))
	must.Nil(err)
	err = w.Close()
	must.Nil(err)
	
	ft,r,err := cnt.NextReader()
	must.Nil(err)
	should.Equal(BINARY, ft)
	b, err := ioutil.ReadAll(r)
	must.Nil(err)
	should.Equal([]byte{1,2,3,4},b)
	err = r.Close()
	must.Nil(err)

	cnt.Close()

	wg.Wait()
}