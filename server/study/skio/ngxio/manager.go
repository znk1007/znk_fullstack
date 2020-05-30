package ngxio

import (
	"strconv"
	"sync"
	"sync/atomic"
)

//SessionIDGenerator generates new session id.
type SessionIDGenerator interface {
	NewID() string
}

type defaultIDGenrator struct {
	nextID uint64
}

func (g *defaultIDGenrator) NewID() string {
	id := atomic.AddUint64(&g.nextID, 1)
	return strconv.FormatUint(id, 36)
}

type manager struct {
	SessionIDGenerator
	s      map[string]*session
	locker sync.RWMutex
}

func newManager(g SessionIDGenerator) *manager {
	if g == nil {
		g = &defaultIDGenrator{}
	}
	return &manager{
		SessionIDGenerator: g,
		s:                  make(map[string]*session),
	}
}

func (m *manager) Add(s *session) {
	m.locker.RLock()
	defer m.locker.RUnlock()
	m.s[s.ID()] = s
}

func (m *manager) Get(sid string) *session {
	m.locker.RLock()
	defer m.locker.RUnlock()
	return m.s[sid]
}

func (m *manager) Remove(sid string) {
	m.locker.Lock()
	defer m.locker.Unlock()
	if _, ok := m.s[sid]; !ok {
		return
	}
	delete(m.s, sid)
}
