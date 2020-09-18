package anagram

import (
	"sort"
	"strings"
)

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	r1 := append([]rune{}, []rune(s1)...)
	r2 := append([]rune{}, []rune(s2)...)

	sort.Slice(r1, func(i int, j int) bool { return r1[i] > r1[j] })
	sort.Slice(r2, func(i int, j int) bool { return r2[i] > r2[j] })

	return strings.Compare(string(r1), string(r2)) == 0
}
