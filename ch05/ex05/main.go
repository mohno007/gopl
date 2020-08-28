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
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func countWords(text string) int {
	reader := strings.NewReader(text)
	input := bufio.NewScanner(reader)
	input.Split(bufio.ScanWords)

	count := 0
	for input.Scan() {
		count++
	}
	return count
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if len(text) > 0 {
			words += countWords(text)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "img" {
			images++
		} else if (c.Type == html.ElementNode && c.Data != "script" && c.Data != "style") || c.Type != html.ElementNode {
			cWords, cImages := countWordsAndImages(c)
			words += cWords
			images += cImages
		}
	}
	return
}

//!+
func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "count_word_and_images: %v\n", err)
			continue
		}
		fmt.Printf("words: %v\n", words)
		fmt.Printf("images: %v\n", images)
	}
}

func CountWordAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		err = fmt.Errorf("getting %s: %s", url, resp.Status)
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

//!-
