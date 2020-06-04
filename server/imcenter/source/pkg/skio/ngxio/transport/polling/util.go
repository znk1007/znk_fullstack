package polling

import (
	"errors"
	"mime"
	"strings"
)

func mimeSupportBinary(m string) (bool, error) {
	t, p, e := mime.ParseMediaType(m)
	if e != nil {
		return false, e
	}
	switch t {
	case "application/octet-stream":
		return true, nil
	case "text/plain":
		charset := strings.ToLower(p["charset"])
		if charset != "utf-8" {
			return false, errors.New("invalid charset")
		}
		return false, nil
	}
	return false, errors.New("invalid content-type")
}

//Addr is address of the polling
type Addr struct {
	Host string
}

//Network is network of the polling
func (a Addr) Network() string {
	return "tcp"
}

func (a Addr) String() string {
	return a.Host
}
