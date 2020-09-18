package popcnt

import (
	"crypto/sha256"
	"testing"
)

func TestSha256Diff(t *testing.T) {
	tc := [...]struct {
		input1   [32]uint8
		input2   [32]uint8
		expected int
	}{
		{sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X")), 125},
		{sha256.Sum256([]byte("y")), sha256.Sum256([]byte("Y")), 119},
		{sha256.Sum256([]byte("z")), sha256.Sum256([]byte("Z")), 119},
	}

	for _, c := range tc {
		actual := sha256Diff(c.input1, c.input2)
		if actual != c.expected {
			t.Errorf("given(%x, %x): expected %d, actual %d", c.input1, c.input2, c.expected, actual)
		}
	}
}
