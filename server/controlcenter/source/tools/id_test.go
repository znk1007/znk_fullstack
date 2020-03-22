package tools

import "testing"

func TestGenerateID(t *testing.T) {

}

func Benchmark_GenerateID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateID()
	}
}
