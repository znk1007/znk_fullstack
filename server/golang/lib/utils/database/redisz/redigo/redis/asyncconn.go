package redis

import (
	"errors"
	"time"
)

// Resultz 连接结果
type Resultz struct {
	result interface{}
	err    error
}

//Requestz 发送do请求
type Requestz struct {
	cmd  string
	args []interface{}
	c    chan *Resultz
}

// Replyz 请求响应
type Replyz struct {
	cmd string
	c   chan *Resultz
}

// asyncResultz 异步请求结果
type asyncResultz struct {
	c chan *Resultz
}

// asyncConn异步连接
type asyncConn struct {
	*conn
	t            time.Time
	reqChan      chan *Requestz
	repChan      chan *Replyz
	closeReqChan chan bool
	closeRepChan chan bool
	closed       bool
}

// AsyncDailTimeout 异步拨号连接redis服务器，可设置超时
func AsyncDailTimeout(network, address string, connectTimeout, readTimeout, writeTimeout time.Duration) (AsyncConn, error) {
	return AsyncDial(
		network,
		address,
		DialConnectTimeout(connectTimeout),
		DialReadTimeout(readTimeout),
		DialWriteTimeout(writeTimeout),
	)
}

// AsyncDial 异步拨号连接redis服务器
func AsyncDial(network, address string, options ...DialOption) (AsyncConn, error) {
	tmp, err := Dial(network, address, options...)
	if err != nil {
		return nil, err
	}
	return getAsyncConn(tmp.(*conn))
}

// AsyncDialURL 异步拨号连接redis服务器
func AsyncDialURL(rawurl string, options ...DialOption) (AsyncConn, error) {
	tmp, err := DialURL(rawurl, options...)
	if err != nil {
		return nil, err
	}
	return getAsyncConn(tmp.(*conn))
}

// getAsyncConn 获取异步连接
func getAsyncConn(conn *conn) (AsyncConn, error) {
	c := &asyncConn{
		conn:         conn,
		reqChan:      make(chan *Requestz, 1000),
		repChan:      make(chan *Replyz, 1000),
		closeReqChan: make(chan bool),
		closeRepChan: make(chan bool),
	}
	go c.doRequest()
	go c.doReply()
	return c, nil
}

// Do操作redis数据库
func (c *asyncConn) Do(cmd string, args ...interface{}) (interface{}, error) {
	if cmd == "" {
		return nil, errors.New("empty command is not allow")
	}
	retChan := make(chan *Resultz, 2)
	c.reqChan <- &Requestz{
		cmd:  cmd,
		args: args,
		c:    retChan,
	}
	ret := <-retChan
	if ret.err != nil {
		return ret.result, ret.err
	}
	ret = <-retChan
	return ret.result, ret.err
}

// AsyncDo 异步操作数据库
func (c *asyncConn) AsyncDo(cmd string, args ...interface{}) (AsyncResultz, error) {
	if cmd == "" {
		return nil, errors.New("empty command is not allow")
	}
	retChan := make(chan *Resultz, 2)
	c.reqChan <- &Requestz{
		cmd:  cmd,
		args: args,
		c:    retChan,
	}
	return &asyncResultz{
		c: retChan,
	}, nil
}

// Get 获取结果
func (r *asyncResultz) Get() (interface{}, error) {
	send := <-r.c
	if send.err != nil {
		return send.result, send.err
	}
	recv := <-r.c
	return recv.result, recv.err
}

// Close 关闭数据库连接
func (c asyncConn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return nil
	}
	c.closed = true
	go func() {
		time.Sleep(10 * time.Minute)
		c.closeReqChan <- true
	}()
	c.err = errors.New("closed")
	return c.conn.conn.Close()
}

// doRequest 异步请求
func (c *asyncConn) doRequest() {
	for {
		select {
		case <-c.closeReqChan:
			close(c.reqChan)
			c.closeRepChan <- true
			return
		case req := <-c.reqChan:
			for i, length := 0, len(c.reqChan); ; {
				if c.writeTimeout != 0 {
					c.conn.conn.SetWriteDeadline(time.Now().Add(c.writeTimeout))
				}
				if err := c.writeCommand(req.cmd, req.args); err != nil {
					req.c <- &Resultz{nil, err}
					c.fatal(err)
					break
				}
				req.c <- &Resultz{nil, nil}
				c.repChan <- &Replyz{
					cmd: req.cmd,
					c:   req.c,
				}
				if i++; i > length {
					break
				}
				req = <-c.reqChan
			}
		}
		if err := c.bw.Flush(); err != nil {
			c.fatal(err)
			continue
		}
	}
}

// doReply 请求响应
func (c *asyncConn) doReply() {
	for {
		select {
		case <-c.closeRepChan:
			close(c.repChan)
			return
		case rep := <-c.repChan:
			if c.readTimeout != 0 {
				c.conn.conn.SetReadDeadline(time.Now().Add(c.readTimeout))
			}
			reply, err := c.readReply()
			if err != nil {
				rep.c <- &Resultz{nil, err}
				c.fatal(err)
				continue
			} else {
				c.t = nowFunc()
			}
			if e, ok := reply.(Error); ok {
				err = e
			}
			rep.c <- &Resultz{reply, err}
		}
	}
}
