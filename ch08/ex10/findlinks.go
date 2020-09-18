// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/mohno007/gopl/ch08/ex10/links"
)

type WorkList struct {
	list  []string
	depth int64
}

type Link struct {
	url   string
	depth int64
}

var cancel = make(chan struct{})

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url, cancel)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

var maxDepth = flag.Int64("depth", math.MaxInt64, "max depth for searching links")

func cancelled() bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}

//!+
func main() {
	flag.Parse()

	worklist := make(chan WorkList) // lists of URLs, may have duplicates
	var n int = 0

	// Add command-line arguments to worklist.
	n++
	go func() {
		worklist <- WorkList{list: flag.Args(), depth: *maxDepth}
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth < 0 {
			continue
		}
		for _, link := range list.list {
			if cancelled() {
				return
			}
			if seen[link] {
				continue
			}
			seen[link] = true
			if list.depth-1 == 0 {
				continue
			}
			n++
			go func(l string, d int64) {
				foundLinks := crawl(l)
				worklist <- WorkList{list: foundLinks, depth: d - 1}
			}(link, list.depth)
		}
	}

}

//!-
