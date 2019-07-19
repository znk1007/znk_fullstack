package common

import (
	"errors"
	"fmt"
	"net"
)

// Error 错误接口
type Error interface {
	Error() string
	Temporary() bool
}

// Errs socket相关错误
type Errs struct {
	URL       string
	Operation string
	Err       error
	IsNet     bool
}

// NewErr 创建错误
func NewErr(url, operation string, err error) error {
	return &Errs{
		URL:       url,
		Operation: operation,
		Err:       err,
	}
}

func (e *Errs) Error() string {
	if e.URL == "" {
		return fmt.Sprintf("%s:%s", e.Operation, e.Err.Error())
	}
	return fmt.Sprintf("%s %s:%s", e.Operation, e.URL, e.Err.Error())
}

// Timeout 超时错误
func (e *Errs) Timeout() bool {
	if e.IsNet {
		if r, ok := e.Err.(net.Error); ok {
			return r.Timeout()
		}
	}
	return false
}

// Temporary 临时错误
func (e *Errs) Temporary() bool {
	if e.IsNet {
		if err, ok := e.Err.(net.Error); ok {
			return err.Temporary()
		}
	}
	if err, ok := e.Err.(Error); ok {
		return err.Temporary()
	}
	return false
}

// RetryError 重试错误
type RetryError struct {
	err string
}

func (e RetryError) Error() string {
	return e.err
}

// Temporary 重试临时错误
func (e RetryError) Temporary() bool {
	return true
}

var (
	// ErrPaused 暂停错误
	ErrPaused = RetryError{"paused"}
	// ErrTimeout 超时
	ErrTimeout = RetryError{"timeout"}
	// ErrInvalidPayload 无效负载
	ErrInvalidPayload = errors.New("invalid payload")
	// ErrDrain 无效输出
	ErrDrain = errors.New("drain")
	// ErrOverlap 重叠错误
	ErrOverlap = errors.New("overlap")
)
