package payload

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

	p := New(true)
	p.Pause()
	p.Resume()
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for _, test := range tests {
			if len(test.packets) != 1 {
				continue
			}
			r := bytes.NewReader(test.data)
			err := p.FeedIn(r, test.supportBinary)
			must.Nil(err)
		}
	}()
	for _, test := range tests {
		if len(test.packets) != 1 {
			continue
		}
		p.SetReadDeadline(time.Now().Add(time.Second / 10))
		ft, pt, r, err := p.NextReader()
		must.Nil(err)
		should.Equal(test.packets[0].ft, ft)
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
