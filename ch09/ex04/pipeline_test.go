package pipeline

import (
	"testing"
)

func BenchmarkPipeline(b *testing.B) {
	length := uint(2 << 1)
	expected := 1

	for i := 0; i < b.N; i++ {
		done := make(chan struct{})
		b.StopTimer()
		in, out := makePipeline(length, done)
		b.StartTimer()
		in <- expected
		v := <-out
		if v != expected {
			b.Fatalf("expected %v, got %v", expected, v)
		}
		b.StopTimer()
		close(done)
		b.StartTimer()
	}
}
