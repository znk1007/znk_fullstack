package core

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPausingTrigger(t *testing.T) {
	should := assert.New(t)
	pp := newPayloadPauser()
	ok := pp.Working()
	should.True(ok)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ok := pp.Pause()
		should.True(ok)
		defer pp.Resume()
	}()
	select {
	case <-pp.PausingTrigger():
	case <-time.After(time.Second / 10):
		should.True(false, "should not run here")
	}
	select {
	case <-pp.PausedTrigger():
		should.True(false, "should not run here")
	case <-time.After(time.Second / 10):
	}

	go func() {
		time.Sleep(time.Second / 10)
		pp.Done()
	}()
	select {
	case <-pp.PausedTrigger():
	case <-time.After(time.Second):
		should.True(false, "should not run here")
	}
	wg.Wait()
	select {
	case <-pp.PausingTrigger():
		should.True(false, "should not run here")
	case <-pp.PausedTrigger():
		should.True(false, "should not run here")
	case <-time.After(time.Second / 10):
	}
}

func TestPayloadPauserOnlyOnce(t *testing.T) {
	should := assert.New(t)
	pp := newPayloadPauser()
	s := make(chan int)
	go func() {
		ok := pp.Pause()
		should.True(ok)
		defer pp.Resume()
		s <- 1
		<-s
	}()
}

func TestPauseAfterResume(t *testing.T) {
	should := assert.New(t)
	pp := newPayloadPauser()
	ok := pp.Pause()
	should.True(ok)
	pp.Resume()
	ok = pp.Pause()
	should.True(ok)
	pp.Resume()
}

func TestPauseMultiplyResumeOnce(t *testing.T) {
	should := assert.New(t)
	pp := newPayloadPauser()
	ok := pp.Pause()
	should.True(ok)
	for idx := 0; idx < 10; idx++ {
		ok = pp.Pause()
		should.False(ok)
	}
	pp.Resume()
}

func TestPauserConcurrencyWorkingDone(t *testing.T) {
	pp := newPayloadPauser()
	wg := sync.WaitGroup{}
	f := func() {
		defer wg.Done()
		should := assert.New(t)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Microsecond * time.Duration(r.Intn(100)))
		ok := pp.Working()
		should.True(ok)
		defer pp.Done()
		time.Sleep(time.Microsecond * time.Duration(r.Intn(100)))
	}
	max := 100
	wg.Add(max)
	for idx := 0; idx < max; idx++ {
		go f()
	}
	wg.Wait()
}

func TestPauserCanWorkingDurationPauseWaiting(t *testing.T) {
	should := assert.New(t)
	pp := newPayloadPauser()
	wg := sync.WaitGroup{}
	ok := pp.Working()
	should.True(ok)
	wg.Add(1)
	go func() {
		defer wg.Done()
		p := pp.Pause()
		should.True(p)
		defer pp.Resume()
	}()
	<-pp.PausingTrigger()
	ok = pp.Working()
	should.True(ok)

	pp.Done()
	pp.Done()
	wg.Wait()
}

func TestPauserPauseWhenAllDone(t *testing.T) {
	should := assert.New(t)
	pp := newPayloadPauser()
	n := 10
	for idx := 0; idx < n; idx++ {
		ok := pp.Working()
		should.True(ok)
	}
	for idx := 0; idx < n; idx++ {
		pp.Done()
	}
	ok := pp.Pause()
	should.True(ok)

	ok = pp.Pause()
	should.False(ok)
	pp.Resume()
}

func TestPauserOnlyOncePauseAfterWaiting(t *testing.T) {
	should := assert.New(t)
	pp := newPayloadPauser()
	count := int64(0)
	wg := sync.WaitGroup{}

	ok := pp.Working()
	should.True(ok)
	wg.Add(10)
	for idx := 0; idx < 10; idx++ {
		go func() {
			defer wg.Done()
			ok := pp.Pause()
			if ok {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	time.Sleep(time.Second / 10)
	pp.Done()
	wg.Wait()
	should.Equal(int64(1), count)
	pp.Resume()
}

func TestPauserCannotWorkingAffterPause(t *testing.T) {
	should := assert.New(t)
	pp := newPayloadPauser()

	ok := pp.Pause()
	should.True(ok)
	defer pp.Resume()

	ok = pp.Working()
	should.False(ok)
	pp.Done()
}

func TestPauserRandom(t *testing.T) {
	pp := newPayloadPauser()
	wg := sync.WaitGroup{}
	n := 100
	f := func() {
		defer wg.Done()
		should := assert.New(t)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Millisecond * time.Duration(r.Intn(n)))
		ok := pp.Working()
		should.True(ok)
		defer pp.Done()
		time.Sleep(time.Millisecond * time.Duration(r.Intn(n)))
	}
	max := 1000
	wg.Add(max)
	for idx := 0; idx < max; idx++ {
		go f()
	}
	should := assert.New(t)
	ok := pp.Working()
	should.True(ok)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * time.Duration(n/2))
		pp.Done()
	}()
	start := time.Now()
	ok = pp.Pause()
	end := time.Now()
	should.True(ok)
	should.True(end.Sub(start) > time.Millisecond)
	wg.Wait()
}
