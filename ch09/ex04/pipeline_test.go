package pipeline

import (
	"testing"
)

func BenchmarkPipeline(b *testing.B) {
	length := uint(2 << 10)
	expected := 1

	for i := 0; i < b.N; i++ {
		in, out := makePipeline(length)
		in <- expected
		v := <-out
		if v != expected {
			b.Fatalf("expected %v, got %v", expected, v)
		}
	}
}
