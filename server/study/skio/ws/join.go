package ws

import (
	"io"
	"strings"
)

type joinReader struct {
	c    *Conn
	term string
	r    io.Reader
}

func (jr *joinReader) Read(p []byte) (int, error) {
	if jr.r == nil {
		var err error
		_, jr.r, err = jr.c.NextReader()
		if err != nil {
			return 0, err
		}
		if len(jr.term) != 0 {
			jr.r = io.MultiReader(jr.r, strings.NewReader(jr.term))
		}
	}
	n, err := jr.r.Read(p)
	if err == io.EOF {
		err = nil
		jr.r = nil
	}
	return n, err
}

//JoinMessages concatenates received messages to create a single io.Reader.
//The string term is appended to each message.
//The returned reader does not support concurrent calls to the Read method.
func JoinMessages(c *Conn, term string) io.Reader {
	return &joinReader{c: c, term: term}
}
