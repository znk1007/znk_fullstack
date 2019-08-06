package core

import (
	"sync"

	"github.com/znk_fullstack/golang/lib/utils/socket/socketlib/protos/pbs"
)

type payloadPauser struct {
	l       sync.Mutex
	c       *sync.Cond
	worker  int
	pausing chan struct{}
	paused  chan struct{}
	status  pbs.PauserStatus
}

func newPayloadPauser() *payloadPauser {
	ret := &payloadPauser{
		pausing: make(chan struct{}),
		paused:  make(chan struct{}),
		status:  pbs.PauserStatus_normal,
	}
	ret.c = sync.NewCond(&ret.l)
	return ret
}

func (pp *payloadPauser) Pause() bool {
	pp.l.Lock()
	defer pp.l.Unlock()
	switch pp.status {
	case pbs.PauserStatus_paused:
		return false
	case pbs.PauserStatus_normal:
		close(pp.pausing)
		pp.status = pbs.PauserStatus_pausing
	}
	for pp.worker != 0 {
		pp.c.Wait()
	}
	if pp.status == pbs.PauserStatus_paused {
		return false
	}
	close(pp.paused)
	pp.status = pbs.PauserStatus_paused
	pp.c.Broadcast()
	return true
}

func (pp *payloadPauser) Resume() {
	pp.l.Lock()
	defer pp.l.Unlock()
	pp.status = pbs.PauserStatus_normal
	pp.paused = make(chan struct{})
	pp.pausing = make(chan struct{})
}

func (pp *payloadPauser) Working() bool {
	pp.l.Lock()
	defer pp.l.Unlock()
	if pp.status == pbs.PauserStatus_paused {
		return false
	}
	pp.worker++
	return true
}

func (pp *payloadPauser) Done() {
	pp.l.Lock()
	defer pp.l.Unlock()
	if pp.status == pbs.PauserStatus_paused || pp.worker == 0 {
		return
	}
	pp.worker--
	pp.c.Broadcast()
}

func (pp *payloadPauser) PausingTrigger() <-chan struct{} {
	pp.l.Lock()
	defer pp.l.Unlock()
	return pp.pausing
}

func (pp *payloadPauser) PausedTrigger() <-chan struct{} {
	pp.l.Lock()
	defer pp.l.Unlock()
	return pp.paused
}
