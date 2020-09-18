package anagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tc := [...]struct {
		expected bool
		s1       string
		s2       string
	}{
		{true, "tac", "cat"},
		{true, "cinerama", "american"},
		{true, "canoe", "ocean"},
		{true, "とけい", "けいと"},
		{true, "こうたい", "たいこう"},
		{false, "t", "tt"},
		{false, "いいたい", "たたいた"},
	}

	for _, c := range tc {
		result := isAnagram(c.s1, c.s2)

		if result != c.expected {
			t.Fatalf("%s %s %v", c.s1, c.s2, c.expected)
		}
	}
}
