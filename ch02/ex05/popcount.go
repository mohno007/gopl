package popcount

// #include <x86intrin.h>
//
// #cgo CFLAGS: -march=native -O3
import "C"

// 8bit整数をすべて包含する
var pc [256]byte

func init() {
	for i := range pc {
		// iの半分の個数は計算済みのはず
		// 0bXXXY なら 0bXXXは計算済み、0bYのみ求めれば良い
		pc[i] = pc[i>>1] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountWithLoop(x uint64) int {
	result := 0

	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}

	return result
}

func PopCountWithBitShift(x uint64) int {
	result := 0

	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			result++
		}
		x >>= 1
	}

	return result
}

func PopCountWithBitAnd(x uint64) int {
	result := 0

	for x != 0 {
		x &= x - 1
		result++
	}

	return result
}

func PopCountNative(x uint64) int {
	return int(C._popcnt64(C.longlong(x)))
}

// https://qiita.com/ocxtal/items/01c46b15cb1f2e656887#popcnt
// https://en.wikipedia.org/wiki/Hamming_weight#Efficient_implementation
func PopCountHamming(x uint64) int {
	x = (x & 0x5555555555555555) + ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x & 0x0f0f0f0f0f0f0f0f) + ((x >> 4) & 0x0f0f0f0f0f0f0f0f)
	x = (x & 0x00ff00ff00ff00ff) + ((x >> 8) & 0x00ff00ff00ff00ff)
	x = (x & 0x0000ffff0000ffff) + ((x >> 16) & 0x0000ffff0000ffff)
	x = (x & 0x00000000ffffffff) + ((x >> 32) & 0x00000000ffffffff)
	return int(x)
}
