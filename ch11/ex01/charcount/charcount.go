package charcount

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

// CharCountResult represents the
type CharCountResult struct {
	// counts of Unicode characters
	Counts map[rune]int
	// count of lengths of UTF-8 encodings
	Utflen [utf8.UTFMax + 1]int
	// count of invalid UTF-8 characters
	Invalid int
}

// CharCount counts the number of unicode characters in Reader
func CharCount(in bufio.Reader) (*CharCountResult, error) {
	result := CharCountResult{
		Counts:  make(map[rune]int),
		Invalid: 0,
	}

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("charcount: %v", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			result.Invalid++
			continue
		}
		result.Counts[r]++
		result.Utflen[n]++
	}

	return &result, nil
}
