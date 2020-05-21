package ngxio

import (
	"bytes"
	"time"
)

var chars = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_")

//timestamp returns a string based on different nano time.
func timestamp() string {
	now := time.Now().UnixNano()
	ret := make([]byte, 0, 16)
	for now > 0 {
		ret = append(ret, chars[int(now%int64(len(chars)))])
		now = now / int64(len(chars))
	}
	return string(ret)
}

//writeBinaryLen write websocket binary data with len
func writeBinaryLen(l int64, w *bytes.Buffer) error {
	//write head
	if l <= 0 {
		if err := w.WriteByte(0x00); err != nil {
			return err
		}
		if err := w.WriteByte(0xff); err != nil {
			return err
		}
		return nil
	}
	max := int64(1)
	for n := l / 10; n > 0; n /= 10 {
		max *= 10
	}
	for max > 0 {
		n := l / max
		if err := w.WriteByte(byte(n)); err != nil {
			return err
		}
		l -= n * max
		max /= 10
	}
	return w.WriteByte(0xff)
}

//writeTextLen write websocket text data with length
func writeTextLen(l int64, w *bytes.Buffer) error {
	if l <= 0 {
		if err := w.WriteByte('0'); err != nil {
			return err
		}
		if err := w.WriteByte(':'); err != nil {
			return err
		}
		return nil
	}
	max := int64(1)
	for n := l / 10; n > 0; n /= 10 {
		max *= 10
	}
	for max > 0 {
		n := l / max
		if err := w.WriteByte(byte(n) + '0'); err != nil {
			return err
		}
		l -= n * max
		max /= 10
	}
	return w.WriteByte(':')
}

//readBinaryLen read websocket binary data with length
func readBinaryLen(r byteReader) (int64, error) {
	ret := int64(0)
	for {

	}
}
