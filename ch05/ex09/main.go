package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func scanPlaceholder(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start, end := -1, -1

	if len(data) == 0 {
		return 0, nil, nil
	}

	// Start with Placeholder
	for i, w := 0, 0; i < len(data); i += w {
		runeValue, width := utf8.DecodeRune(data[i:])

		// Error
		if runeValue == utf8.RuneError {
			return 0, nil, fmt.Errorf("input is invalid UTF-8 string: %v", data)
		}

		// End Placeholder
		if start != -1 && !unicode.IsLetter(runeValue) {
			end = i + width - 1
			break
		}

		// Start Placeholder
		if runeValue == '$' {
			// Placeholder not start at position 0
			if i != 0 {
				advance = i + width - 1
				token = data[0:advance]
				return
			}
			start = i
		}
		w = width
	}

	// any placeholder not found
	if start == -1 {
		return len(data), data, nil
	}

	// placeholder found, but not ended
	if end == -1 {
		if !atEOF {
			// request more bytes
			return 0, nil, nil
		}
		end = len(data)
	}

	advance = end
	token = data[start:advance]
	err = nil
	return
}

func expand(s string, f func(string) string) string {
	r := strings.NewReader(s)
	sc := bufio.NewScanner(r)
	sc.Split(scanPlaceholder)

	sb := strings.Builder{}
	for sc.Scan() {
		text := sc.Text()
		if len(text) > 0 && text[0] == '$' {
			sb.WriteString(f(text[1:]))
		} else {
			sb.WriteString(text)
		}
	}
	return sb.String()
}

func main() {
	m := map[string]string{
		"a": "x",
		"b": "y",
		"c": "z",
	}

	result := expand(
		" a $a b $b c $c ",
		func(s string) string { return m[s] },
	)
	fmt.Printf("%s\n", result)
}
