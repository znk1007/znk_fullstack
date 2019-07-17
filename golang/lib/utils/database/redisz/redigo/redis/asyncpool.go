package redis

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var errorCompatibility = errors.New("Should use AsyncDo func")

// 异步池连接
type asyncPoolConn struct {
	p *AsyncPool
	c AsyncConn
}

// 连接池错误
func (pc *asyncPoolConn) Err() error {
	return pc.c.Err()
}

func (pc *asyncPoolConn) Close() error {
	return nil
}

func (pc *asyncPoolConn) AsyncDo(commandName string, args ...interface{}) (resultz AsyncResultz, err error) {
	return pc.c.AsyncDo(commandName, args...)
}

// 以连接池操作数据库
func (pc *asyncPoolConn) Do(commandName string, args ...interface{}) (interface{}, error) {
	if pc.p.MaxDoCount != 0 {
		if atomic.AddInt32(&pc.p.doCount, 1) > int32(pc.p.MaxDoCount) {
			atomic.AddInt32(&pc.p.doCount, -1)
			return nil, ErrPoolExhausted
		}
		defer func() {
			atomic.AddInt32(&pc.p.doCount, -1)
		}()
	}
	return pc.c.Do(commandName, args...)
}

func (pc *asyncPoolConn) Send(commandName string, args ...interface{}) error {
	return errorCompatibility
}

func (pc *asyncPoolConn) Flush() error {
	return errorCompatibility
}

func (pc *asyncPoolConn) Receive() (interface{}, error) {
	return nil, errConnClosed
}

func (ec errorConnection) AsyncDo(string, ...interface{}) (AsyncResultz, error) {
	return nil, ec.err
}

// AsyncPool 异步连接池
type AsyncPool struct {
	// 拨号连接
	Dail func() (AsyncConn, error)
	// 检查连接是否健壮
	TestOnBorrow func(c AsyncConn, t time.Time) error
	// 控制Get()方法协程数，为0则不限制
	MaxGetCount int
	// 控制Do()方法协程数，为0则不限制
	MaxDoCount int
	// 异步连接
	c *asyncPoolConn
	// 互斥锁
	mu sync.Mutex
	// 异步执行条件
	cond *sync.Cond
	// 获取连接数
	getCount int
	// Do执行次数
	doCount int32
	// 关闭
	closed bool
	// 块执行
	blocking bool
}

// NewAsyncPool 初始化异步连接池
func NewAsyncPool(newFn func() (AsyncConn, error), testFn func(AsyncConn, time.Time) error) *AsyncPool {
	return &AsyncPool{
		Dail:         newFn,
		TestOnBorrow: testFn,
	}
}

// Get 获取连接池中的连接
func (p *AsyncPool) Get() AsyncConn {
	p.mu.Lock()
	if p.cond == nil {
		p.cond = sync.NewCond(&p.mu)
	}
	p.getCount++
	if p.MaxGetCount != 0 && p.getCount > p.MaxGetCount {
		p.getCount--
		p.mu.Unlock()
		return errorConnection{ErrPoolExhausted}
	}
	var pc AsyncConn
	for {
		if p.closed {
			p.getCount--
			p.mu.Unlock()
			return errorConnection{errPoolClosed}
		}
		if p.blocking {
			p.cond.Wait()
			continue
		}
		if p.c != nil && p.c.Err() == nil {
			if test := p.TestOnBorrow; test != nil {
				p.blocking = true
				ic := p.c.c.(*asyncConn)
				p.mu.Unlock()

				err := test(p.c, ic.t)

				p.mu.Lock()
				p.blocking = false
				if err == nil {
					pc = p.c
					p.getCount--
					p.cond.Signal()
					p.mu.Unlock()
					return pc
				}
			} else {
				pc = p.c
				p.getCount--
				p.cond.Signal()
				p.mu.Unlock()
				return pc
			}
		}
		if p.c != nil {
			p.c.c.Close()
		}
		p.blocking = true
		p.mu.Unlock()

		c, err := p.Dail()

		p.mu.Lock()
		p.blocking = false

		if err != nil {
			p.getCount--
			p.cond.Signal()
			p.mu.Unlock()
			return errorConnection{err}
		}
		p.c = &asyncPoolConn{
			p: p,
			c: c,
		}
		pc := p.c
		p.getCount--
		p.cond.Signal()
		p.mu.Unlock()
		return pc
	}
}

// ActiveCount 连接池活跃数
func (p *AsyncPool) ActiveCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.c != nil && p.c.Err() == nil {
		return 1
	}
	return 0
}

// IdleCount 空闲连接池数
func (p *AsyncPool) IdleCount() int {
	return 0
}

// Close 关闭连接池
func (p *AsyncPool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return nil
	}
	p.closed = true
	if p.cond != nil {
		p.cond.Broadcast()
	}
	err := p.c.c.Close()
	p.c = nil
	return err
}
