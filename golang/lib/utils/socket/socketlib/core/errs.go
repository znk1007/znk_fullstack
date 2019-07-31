package socket

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

// Error payload error interface
type Error interface {
	Error() string
	Temporary() bool
}

// retryError error for retry to connect
type retryError struct {
	err string
}

func (r *retryError) Error() string {
	return r.err
}

func (r *retryError) Temporary() bool {
	return true
}

// Err error for socketlib
type Err struct {
	URL       string
	Operation string
	E         error
}

// NewErr new *Err
func NewErr(url, operation string, err error) *Err {
	return &Err{
		URL:       url,
		Operation: operation,
		E:         err,
	}
}

func (e *Err) Error() string {
	if e.URL == "" {
		return fmt.Sprintf("%s: %s", e.Operation, e.E.Error())
	}
	return fmt.Sprintf("%s with %s: %s", e.URL, e.Operation, e.E.Error())
}

// Timeout if is timeout error
func (e *Err) Timeout() bool {
	if r, ok := e.E.(net.Error); ok {
		return r.Timeout()
	}
	return false
}

// Temporary the err is temporay or not
func (e *Err) Temporary() bool {
	if r, ok := e.E.(net.Error); ok {
		return r.Temporary()
	}
	return false
}
