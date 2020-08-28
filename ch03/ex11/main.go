// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func commaBuffer(s string) string {
	var buf bytes.Buffer
	n := len(s)
	i := 0

	if n == 0 {
		return ""
	}

	// sign
	if s[i] == '+' || s[i] == '-' {
		buf.WriteByte(s[i])
		i++
	}

	// integer part
	ipoint := strings.Index(s, ".")
	digits := (ipoint - i) % 3
	if digits == 0 {
		digits = 3
	}

	for i < ipoint {
		buf.WriteString(s[i : i+digits])
		i += digits
		digits = 3
		if i < ipoint {
			buf.WriteByte(',')
		}
	}

	// fractional part
	if i < n {
		buf.WriteString(s[i:])
	}

	return buf.String()
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-
