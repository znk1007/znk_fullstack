package common

import (
	"testing"
)

func TestNewSocketID(t *testing.T) {
	ch := make(chan ID)
	cnt := 1000000
	for idx := 0; idx < cnt; idx++ {
		go func() {
			id := NewSocketID()
			ch <- id
		}()
	}
	defer close(ch)
	m := make(map[ID]int)
	for idx := 0; idx < cnt; idx++ {
		id := <-ch
		_, ok := m[id]
		if ok {
			t.Errorf("ID: %v is not unique\n", id)
		}
		m[id] = idx
	}
}
