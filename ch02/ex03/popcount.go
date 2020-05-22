package popcount

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
