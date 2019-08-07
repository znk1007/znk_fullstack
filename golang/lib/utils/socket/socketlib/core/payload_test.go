package core

import (
	"bytes"
	"io/ioutil"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPayloadFeedIn(t *testing.T) {
	should := assert.New(t)
	must := require.New(t)

	p := NewPayload(true)
	p.Pause()
	p.Resume()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, test := range payloadTests {
			if len(test.packets) != 1 {
				continue
			}
			r := bytes.NewReader(test.data)
			err := p.FeedIn(r, test.supportBinary)
			must.Nil(err)
		}
	}()
	for _, test := range payloadTests {
		if len(test.packets) != 1 {
			continue
		}
		p.SetReadDeadline(time.Now().Add(time.Second / 10))
		dt, pt, r, err := p.NextReader()
		must.Nil(err)
		t.Log("test packets[0].dt: ", test.packets[0].dt)
		t.Log("dt: ", dt)
		// t.Log("test packets[0].pt: ", test.packets[0].pt)
		// t.Log("pt: ", pt)
		should.Equal(test.packets[0].dt, dt)
		should.Equal(test.packets[0].pt, pt)
		b, err := ioutil.ReadAll(r)
		must.Nil(err)
		err = r.Close()
		must.Nil(err)
		should.Equal(test.packets[0].data, b)
	}
	p.SetReadDeadline(time.Now().Add(time.Second / 10))
	_, _, _, err := p.NextReader()
	should.Equal("read: timeout", err.Error())
	wg.Wait()
}
