package main

var pc [256]uint8

func init() {
	for i := range pc {
		// iの半分の個数は計算済みのはず
		// 0bXXXY なら 0bXXXは計算済み、0bYのみ求めれば良い
		pc[i] = pc[i>>1] + uint8(i&1)
	}
}

func sha256Diff(c1 [32]uint8, c2 [32]uint8) int {
	return popcnt(xor([]uint8(c1[:]), []uint8(c2[:])))
}

func xor(a1 []uint8, a2 []uint8) []uint8 {
	minlen := 0
	if len(a1) < len(a2) {
		minlen = len(a1)
	} else {
		minlen = len(a2)
	}

	result := make([]uint8, minlen)

	for i := range result {
		result[i] = a1[i] ^ a2[i]
	}

	return result
}

func popcnt(a []uint8) int {
	sum := 0
	for _, v := range a {
		sum += int(pc[v])
	}

	return sum
}
