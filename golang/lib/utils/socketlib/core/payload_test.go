package core

import (
	"bytes"
	"io"
	"io/ioutil"
	"sync"
	"testing"
	"time"

	"github.com/znk_fullstack/golang/lib/utils/socketlib/protos/pbs"

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

func TestPayloadFlushOutText(t *testing.T) {
	should := assert.New(t)
	must := require.New(t)
	supportBinary := false
	p := NewPayload(supportBinary)
	p.Pause()
	p.Resume()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		should := assert.New(t)
		must := require.New(t)
		defer wg.Done()
		for _, test := range payloadTests {
			if len(test.packets) != 1 {
				continue
			}
			if test.supportBinary != supportBinary {
				continue
			}
			buf := bytes.NewBuffer(nil)
			err := p.FlushOut(buf)
			must.Nil(err)
			should.Equal(test.data, buf.Bytes())
		}
	}()
	for _, test := range payloadTests {
		if len(test.packets) != 1 {
			continue
		}
		if test.supportBinary != supportBinary {
			continue
		}
		p.SetWriteDeadline(time.Now().Add(time.Second / 10))
		w, err := p.NextWriter(test.packets[0].dt, test.packets[0].pt)
		must.Nil(err)
		_, err = w.Write(test.packets[0].data)
		must.Nil(err)
		err = w.Close()
		must.Nil(err)
	}
	p.SetWriteDeadline(time.Now().Add(time.Second / 10))
	_, err := p.NextWriter(pbs.DataType_binary, pbs.PacketType_open)
	should.Equal("write: timeout", err.Error())
	wg.Wait()
}

func TestPayloadFlushOutBinary(t *testing.T) {
	should := assert.New(t)
	must := require.New(t)
	supportBinary := true
	p := NewPayload(supportBinary)
	p.Pause()
	p.Resume()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		should := assert.New(t)
		must := require.New(t)
		defer wg.Done()
		for _, test := range payloadTests {
			if len(test.packets) != 1 {
				continue
			}
			if test.supportBinary != supportBinary {
				continue
			}
			buf := bytes.NewBuffer(nil)
			err := p.FlushOut(buf)
			must.Nil(err)
			should.Equal(test.data, buf.Bytes())
		}
	}()
	for _, test := range payloadTests {
		if len(test.packets) != 1 {
			continue
		}
		if test.supportBinary != supportBinary {
			continue
		}
		p.SetWriteDeadline(time.Now().Add(time.Second / 10))
		w, err := p.NextWriter(test.packets[0].dt, test.packets[0].pt)
		must.Nil(err)
		_, err = w.Write(test.packets[0].data)
		must.Nil(err)
		err = w.Close()
		must.Nil(err)
	}
	p.SetWriteDeadline(time.Now().Add(time.Second / 10))
	_, err := p.NextWriter(pbs.DataType_binary, pbs.PacketType_open)
	should.Equal("write: timeout", err.Error())
	wg.Wait()
}

func TestPayloadWaitNextClose(t *testing.T) {
	should := assert.New(t)
	p := NewPayload(true)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		should := assert.New(t)
		defer wg.Done()
		_, _, _, err := p.NextReader()
		should.Equal(io.EOF, err)
	}()
	wg.Add(1)
	go func() {
		should := assert.New(t)
		defer wg.Done()
		_, err := p.NextWriter(pbs.DataType_binary, pbs.PacketType_open)
		should.Equal(io.EOF, err)
	}()
	time.Sleep(time.Second / 10)
	err := p.Close()
	should.Nil(err)

	wg.Wait()

	_, _, _, err = p.NextReader()
	should.Equal(io.EOF, err)
	_, err = p.NextWriter(pbs.DataType_binary, pbs.PacketType_open)
	should.Equal(io.EOF, err)
	err = p.FeedIn(bytes.NewReader([]byte("1:0")), false)
	should.Equal(io.EOF, err)
	err = p.FlushOut(ioutil.Discard)
	should.Equal(io.EOF, err)
}

func TestPayloadWaitInOutClose(t *testing.T) {
	should := assert.New(t)
	p := NewPayload(true)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		should := assert.New(t)
		defer wg.Done()
		err := p.FeedIn(bytes.NewBuffer([]byte("1:0")), false)
		should.Equal(io.EOF, err)
	}()
	wg.Add(1)
	go func() {
		should := assert.New(t)
		defer wg.Done()
		err := p.FlushOut(ioutil.Discard)
		should.Equal(io.EOF, err)
	}()
	time.Sleep(time.Second / 10)
	err := p.Close()
	should.Nil(err)
	wg.Wait()

	_, _, _, err = p.NextReader()
	should.Equal(io.EOF, err)
	_, err = p.NextWriter(pbs.DataType_binary, pbs.PacketType_open)
	should.Equal(io.EOF, err)

	err = p.FeedIn(bytes.NewReader([]byte("1:0")), false)
	should.Equal(io.EOF, err)
	err = p.FlushOut(ioutil.Discard)
	should.Equal(io.EOF, err)
}

func TestPayloadPauseClose(t *testing.T) {
	should := assert.New(t)
	p := NewPayload(true)
	p.Pause()
	err := p.Close()
	should.Nil(err)
	_, _, _, err = p.NextReader()
	should.Equal(io.EOF, err)
	_, err = p.NextWriter(pbs.DataType_binary, pbs.PacketType_open)
	should.Equal(io.EOF, err)
	err = p.FeedIn(bytes.NewReader([]byte("1:0")), false)
	should.Equal(io.EOF, err)
	err = p.FlushOut(ioutil.Discard)
	should.Equal(io.EOF, err)
}

func TestPayloadNextPause(t *testing.T) {
	should := assert.New(t)
	p := NewPayload(true)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		should := assert.New(t)
		must := require.New(t)
		defer wg.Done()
		_, _, _, err := p.NextReader()
		op, ok := err.(pError)
		must.True(ok)
		should.True(op.Temporary())
	}()
	wg.Add(1)
	go func() {
		should := assert.New(t)
		must := require.New(t)
		defer wg.Done()
		_, err := p.NextWriter(pbs.DataType_binary, pbs.PacketType_open)
		op, ok := err.(pError)
		t.Log("ok: ", ok)
		t.Log("op: ", op)
		must.True(ok)
		should.True(op.Temporary())
	}()
	time.Sleep(time.Second / 10)
	p.Pause()
	wg.Wait()

	_, _, _, err := p.NextReader()
	op, ok := err.(pError)
	should.True(ok)
	should.True(op.Temporary())
	_, err = p.NextWriter(pbs.DataType_binary, pbs.PacketType_open)
	op, ok = err.(pError)
	should.True(ok)
	should.True(op.Temporary())

	err = p.FeedIn(bytes.NewReader([]byte("1:0")), false)
	op, ok = err.(pError)
	should.True(ok)
	should.True(op.Temporary())

	b := bytes.NewBuffer(nil)
	err = p.FlushOut(b)
	should.Nil(err)
	should.Equal([]byte{0x0, 0x1, 0xff, '6'}, b.Bytes())
}

func TestPayloadInOutPause(t *testing.T) {
	should := assert.New(t)

	p := NewPayload(true)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		must := require.New(t)
		defer wg.Done()
		err := p.FeedIn(bytes.NewReader([]byte("1:0")), false)
		must.Nil(err)
	}()
	wg.Add(1)
	go func() {
		should := assert.New(t)
		must := require.New(t)
		defer wg.Done()

		b := bytes.NewBuffer(nil)
		err := p.FlushOut(b)
		must.Nil(err)
		should.Equal([]byte{0x0, 0x1, 0xff, '6'}, b.Bytes())
	}()

	go func() {
		must := require.New(t)
		time.Sleep(time.Second / 10 * 3)
		_, _, r, err := p.NextReader()
		must.Nil(err)
		defer r.Close()
		io.Copy(ioutil.Discard, r)
	}()

	time.Sleep(time.Second / 10)
	start := time.Now()
	p.Pause()
	end := time.Now()
	should.True(end.Sub(start) >= time.Second/10)

	wg.Wait()

	_, _, _, err := p.NextReader()
	op, ok := err.(pError)
	should.True(ok)
	should.True(op.Temporary())

	_, err = p.NextWriter(pbs.DataType_binary, pbs.PacketType_open)
	op, ok = err.(pError)
	should.True(ok)
	should.True(op.Temporary())

	err = p.FeedIn(bytes.NewReader([]byte("1:0")), false)
	op, ok = err.(pError)
	should.True(ok)
	should.True(op.Temporary())

	b := bytes.NewBuffer(nil)
	err = p.FlushOut(b)
	should.Nil(err)
	should.Equal([]byte{0x0, 0x1, 0xff, '6'}, b.Bytes())
}
