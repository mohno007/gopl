// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	tests := []struct {
		given  []int
		mapset map[int]bool
	}{
		{
			[]int{1, 144, 9},
			map[int]bool{1: true, 144: true, 9: true},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v", tt.given), func(t *testing.T) {
			var intset IntSet
			for _, v := range tt.given {
				intset.Add(v)
			}
			for _, v := range tt.given {
				expected, actual := tt.mapset[v], intset.Has(v)
				if expected != actual {
					t.Errorf(
						"intset.Has(%v) (%v) == mapset[%v] (%v)",
						tt.given,
						actual,
						tt.given,
						expected,
					)
				}
			}
		})
	}
}

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
