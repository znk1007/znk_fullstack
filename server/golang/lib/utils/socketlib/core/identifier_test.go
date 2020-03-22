package core

import (
	"fmt"
	"testing"
)

// test 指令
// go test .

// benchmark 指令
// go test -bench=.
// go test -test.bench=identifier_test.go
func TestNewSocketID(t *testing.T) {
	for idx := 0; idx < 1000000; idx++ {
		id := NewSocketID()
		fmt.Println("id test == ", id.String())
	}
}

func BenchMarkNewSocketID(b *testing.B) {
	for idx := 0; idx < b.N; idx++ {
		id := NewSocketID()
		fmt.Println("id bench == ", id.String())
	}
}
