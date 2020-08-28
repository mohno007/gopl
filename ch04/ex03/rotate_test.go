package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	var tests = []struct {
		given    []int
		expected []int
		amount   int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, []int{4, 5, 0, 1, 2, 3}, 2},
		{[]int{0, 1, 2, 3, 4, 5}, []int{5, 0, 1, 2, 3, 4}, -1},
		{[]int{0, 1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5}, 0},
		{[]int{0, 1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 0}, -1},
		{[]int{0, 1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 0, 1}, -2},
	}
	for _, tt := range tests {
		actual := make([]int, len(tt.given))
		copy(actual, tt.given)
		rotate(actual, tt.amount)
		if equals(actual, tt.expected) {
			t.Errorf("given(%v, %v): expected %v, actual %v", tt.given, tt.amount, tt.expected, actual)
		}
	}
}

func equals(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
