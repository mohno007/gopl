package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	for i := 0; i < 1024; i++ {
		PopCount(uint64(i))
	}
}

func TestPopCountWithLoop(t *testing.T) {
	for i := 0; i < 1024; i++ {
		res := PopCountWithLoop(uint64(i))
		ex := PopCount(uint64(i))
		if res != ex {
			t.Fatalf("invalid value was returned for %d (expected %d, but given %d)", i, ex, res)
		}
	}
}

func TestPopCountWithBitShift(t *testing.T) {
	for i := 0; i < 1024; i++ {
		res := PopCountWithBitShift(uint64(i))
		ex := PopCount(uint64(i))
		if res != ex {
			t.Fatalf("invalid value was returned for %d (expected %d, but given %d)", i, ex, res)
		}
	}
}

func TestPopCountWithBitAnd(t *testing.T) {
	for i := 0; i < 1024; i++ {
		res := PopCountWithBitAnd(uint64(i))
		ex := PopCount(uint64(i))
		if res != ex {
			t.Fatalf("invalid value was returned for %d (expected %d, but given %d)", i, ex, res)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithLoop(uint64(i))
	}
}

func BenchmarkPopCountWithBitShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithBitShift(uint64(i))
	}
}

func BenchmarkPopCountWithBitAnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithBitAnd(uint64(i))
	}
}

func BenchmarkPopCountNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountNative(uint64(i))
	}
}

func BenchmarkPopCountHamming(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountHamming(uint64(i))
	}
}
