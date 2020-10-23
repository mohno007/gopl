// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mohno007/gopl/ch11/ex01/charcount"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	result, err := charcount.CharCount(*in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range result.Counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range result.Utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if result.Invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", result.Invalid)
	}
}

//!-
