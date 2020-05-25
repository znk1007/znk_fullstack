package payload

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPauserTrigger(t *testing.T) {
	should := assert.New(t)

	p := newPauser()

	ok := p.Working()
	should.True(ok)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ok := p.Pause()
		should.True(ok)
		defer p.Resume()
	}()

	select {
	case <-p.PausingTrigger():
	case <-time.After(time.Second / 10):
		should.True(false, "should not run here")
	}

	select {
	case <-p.PausedTrigger():
		should.True(false, "should not run here")
	case <-time.After(time.Second / 10):
	}

	go func() {
		time.Sleep(time.Second / 10)
		p.Done()
	}()

	select {
	case <-p.PausedTrigger():

	case <-time.After(time.Second):
		should.True(false, "should not run here")
	}

	wg.Wait()

	select {
	case <-p.PausingTrigger():
		should.True(false, "should not run here")
	case <-p.PausedTrigger():
		should.True(false, "should not run here")
	case <-time.After(time.Second / 10):
	}
}
