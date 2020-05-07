package skoperr

import "net"

//OpErr error for transport
type OpErr struct {
	URL string
	Op  string
	Err error
}

//NewOpErr 初始化
func NewOpErr(url, op string, err error) error {
	return &OpErr{}
}

func (oe *OpErr) Error() string {
	return oe.Op + " " + oe.URL + ": " + oe.Err.Error()
}

//Timeout returns true if the error is timeout
func (oe *OpErr) Timeout() bool {
	if r, ok := oe.Err.(net.Error); ok {
		return r.Timeout()
	}
	return false
}

//Temporary return true if the error is temporary(临时的)
func (oe *OpErr) Temporary() bool {
	if r, ok := oe.Err.(net.Error); ok {
		return r.Temporary()
	}
	return false
}
