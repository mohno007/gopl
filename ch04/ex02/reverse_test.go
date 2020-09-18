package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		expected [6]int
		given    [6]int
	}{
		{[...]int{6, 5, 4, 3, 2, 1}, [...]int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		reverse(&tt.given)
		if tt.given != tt.expected {
			t.Errorf("expected %v, actual %v", tt.given, tt.given)
		}
	}
}
