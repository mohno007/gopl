// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 125.

// Findlinks2 does an HTTP GET on each URL, parses the
// result as HTML, and prints the links within it.
//
// Usage:
//	findlinks url ...
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(*html.Node) bool) bool {
	if pre != nil {
		shouldContinue := pre(n)
		if !shouldContinue {
			return false
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		shouldContinue := forEachNode(c, pre, post)
		if !shouldContinue {
			return false
		}
	}
	if post != nil {
		shouldContinue := post(n)
		if !shouldContinue {
			return false
		}
	}
	return true
}

//!+
func main() {
	for _, url := range os.Args[1:] {
		err := prettyPrint(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
	}
}

func doesElementHaveId(n *html.Node, id string) bool {
	if n.Type != html.ElementNode {
		return false
	}

	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return true
		}
	}

	return false
}

func findElementById(doc *html.Node, id string) (found *html.Node) {
	find := func(n *html.Node) bool {
		if doesElementHaveId(n, id) {
			found = n
			return false
		}
		return true
	}

	forEachNode(doc, find, find)
	return
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func prettyPrint(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	result := findElementById(doc, "NewReplacer")
	if result != nil {
		fmt.Println(result)
	}

	return nil
}

//!-
