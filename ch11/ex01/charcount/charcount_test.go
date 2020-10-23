package charcount

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		input    []byte
		expected CharCountResult
	}{
		{
			[]byte("aaa"),
			CharCountResult{
				map[rune]int{rune('a'): 3},
				[5]int{0, 3},
				0,
			},
		},
		{
			[]byte("こんばんは"),
			CharCountResult{
				map[rune]int{
					rune('こ'): 1,
					rune('ん'): 2,
					rune('ば'): 1,
					rune('は'): 1,
				},
				[5]int{0, 0, 0, 5},
				0,
			},
		},
	}

	for _, test := range tests {
		in := bufio.NewReader(bytes.NewReader(test.input))
		result, err := CharCount(*in)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(*result, test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}

}
