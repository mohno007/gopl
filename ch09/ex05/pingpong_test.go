package pingpong

import (
	"testing"
	"time"
)

// intを送り合う版 -------------------------------
// 1911664 ops/sec
// 2103958 ops/sec
// 2120440 ops/sec
// 2489213 ops/sec
//
// 190〜250万回 @ Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
// 平均 216万回
//
// struct{}版 -------------------------------
//
// 23066063 pingpong/op
// 22215817 pingpong/op
//
// struct版 GOMAXPROCS=2 -------------------------------
//
// 2708219 ops/sec
// 2230654 ops/sec
// 2293597 ops/sec
func BenchmarkPingpong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		done := make(chan struct{})
		start, result := makePingpong(done)
		start()
		time.Sleep(10 * time.Second)
		close(done)
		count := <-result
		b.ReportMetric(float64(count), "pingpong/op")
	}
}
