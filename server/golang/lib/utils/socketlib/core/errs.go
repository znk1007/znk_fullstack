package core

import (
	"errors"
	"fmt"
	"net"
)

var (
	// errPaused 暂停错误
	errPaused = retryError{"paused"}
	// errTimeout 超时
	errTimeout = retryError{"timeout"}
	// ErrInvalidPayload 无效负载
	errInvalidPayload = errors.New("invalid payload")
	// errDrain 无效输出
	errDrain = errors.New("drain")
	// errOverlap 重叠错误
	errOverlap = errors.New("overlap")
)

// PError payload error interface
type pError interface {
	Error() string
	Temporary() bool
}

// retryError error for retry to connect
type retryError struct {
	err string
}

func (r retryError) Error() string {
	return r.err
}

func (r retryError) Temporary() bool {
	return true
}

// err error for socketlib
type err struct {
	URL       string
	Operation string
	E         error
}

// NewErr new *err
func newErr(url, operation string, e error) *err {
	return &err{
		URL:       url,
		Operation: operation,
		E:         e,
	}
}

func (e *err) Error() string {
	if e.URL == "" {
		return fmt.Sprintf("%s: %s", e.Operation, e.E.Error())
	}
	return fmt.Sprintf("%s with %s: %s", e.URL, e.Operation, e.E.Error())
}

// Timeout if is timeout error
func (e *err) Timeout() bool {
	if r, ok := e.E.(net.Error); ok {
		return r.Timeout()
	}
	return false
}

// Temporary the err is temporay or not
func (e *err) Temporary() bool {
	if r, ok := e.E.(net.Error); ok {
		return r.Temporary()
	} else if oe, ok := e.E.(pError); ok {
		return oe.Temporary()
	}
	return false
}
