package pingpong

import (
	"testing"
	"time"
)

// 1911664.2684887713
// 2103958.0937218512
// 2120440.6386027355
// 2489213.9084448447
//
// 190〜250万回 @ Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
// 平均 216万回
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
