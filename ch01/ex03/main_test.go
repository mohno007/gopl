package ex

import (
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringJoinAppend()
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringJoinJoin()
	}
}
