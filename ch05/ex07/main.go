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

// https://bluesock.org/~willkg/dev/ansi.html
var escapeSequence = map[string]int{
	"R":         0,
	"Reset":     0,
	"Bold":      1,
	"B":         1,
	"Underline": 4,
	"U":         4,
	"Blink":     5,
	"Reverse":   7,
	"NoDisplay": 8,
	"Black":     30,
	"Red":       31,
	"Green":     32,
	"Yellow":    33,
	"Blue":      34,
	"Magenta":   35,
	"Cyan":      36,
	"White":     37,
	"BGBlack":   40,
	"BGRed":     41,
	"BGGreen":   42,
	"BGYellow":  43,
	"BGBlue":    44,
	"BGMagenta": 45,
	"BGCyan":    46,
	"BGWhite":   47,
}

var emptyElements = map[string]bool{
	"area":   true,
	"base":   true,
	"br":     true,
	"col":    true,
	"embed":  true,
	"hr":     true,
	"img":    true,
	"input":  true,
	"link":   true,
	"meta":   true,
	"param":  true,
	"source": true,
	"track":  true,
	"wbr":    true,
}

var depth int = 0

func isEmptyElements(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return false
	}
	if n.FirstChild != nil {
		return false
	}
	_, ok := emptyElements[n.Data]
	return ok
}

func decolate(formats ...string) string {
	if len(formats) == 0 {
		return "\033[0m"
	}

	b := strings.Builder{}
	b.WriteString("\033[0;")
	for i, f := range formats {
		if s, ok := escapeSequence[f]; ok {
			b.WriteString(fmt.Sprintf("%d", s))
		} else {
			continue
		}
		if i+1 != len(formats) {
			b.WriteString(";")
		}
	}
	b.WriteString("m")
	return b.String()
}

func startElement(n *html.Node) string {
	switch n.Type {
	case html.ErrorNode:
		b := strings.Builder{}
		b.WriteString(decolate("BGRed", "Bold"))
		b.WriteString("<!-- ERROR has occurred here: ")
		b.WriteString(n.Data)
		b.WriteString("-->")
		b.WriteString(decolate("R"))
		return b.String()
	case html.TextNode:
		if len(strings.TrimSpace(n.Data)) == 0 {
			return ""
		}
		return n.Data
	case html.DocumentNode:
		b := strings.Builder{}
		b.WriteString(decolate("BGBlue", "Bold"))
		b.WriteString("<!-- DOCUMENT START -->")
		b.WriteString(decolate())
		return b.String()
	case html.ElementNode:
		b := strings.Builder{}
		b.WriteString(decolate("Cyan"))
		b.WriteString("<")
		b.WriteString(decolate("Cyan", "Bold"))
		b.WriteString(n.Data)
		b.WriteString(decolate("Cyan"))
		for _, attr := range n.Attr {
			b.WriteString(" ")
			b.WriteString(attr.Key)
			b.WriteString("=")
			b.WriteString(decolate("Green"))
			b.WriteString("'")
			b.WriteString(attr.Val) // TODO escape quote
			b.WriteString("'")
			b.WriteString(decolate("Cyan"))
		}
		if isEmptyElements(n) {
			b.WriteString(" />")
		} else {
			b.WriteString(">")
		}
		b.WriteString(decolate())
		return b.String()
	case html.CommentNode:
		b := strings.Builder{}
		b.WriteString(decolate("Black", "Bold"))
		b.WriteString("<!--")
		b.WriteString(n.Data)
		b.WriteString(decolate("Black", "Bold"))
		b.WriteString("-->")
		b.WriteString(decolate())
		return b.String()
	case html.DoctypeNode:
		b := strings.Builder{}
		b.WriteString(decolate("Black", "Bold"))
		b.WriteString("<!DOCTYPE ")
		b.WriteString(n.Data)
		b.WriteString(">")
		b.WriteString(decolate())
		return b.String()
	default:
		panic("Unreachable")
	}
}

func endElement(n *html.Node) string {
	switch n.Type {
	case html.ElementNode:
		if !isEmptyElements(n) {
			b := strings.Builder{}
			b.WriteString(decolate("Cyan"))
			b.WriteString("</")
			b.WriteString(decolate("Cyan", "Bold"))
			b.WriteString(n.Data)
			b.WriteString(decolate("Cyan"))
			b.WriteString(">")
			b.WriteString(decolate())
			return b.String()
		}
	case html.DocumentNode:
		b := strings.Builder{}
		b.WriteString(decolate("BGBlue", "Bold"))
		b.WriteString("<!-- DOCUMENT END -->")
		b.WriteString(decolate())
		return b.String()
	}
	return ""
}

func printWithDepth(depth int, i string) {
	WIDTH := 2

	r := strings.NewReader(i)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		width := depth * WIDTH
		for i := 0; i < width; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%s\n", s.Text())
	}
}

// forEachNode appends to links each link found in n, and returns the result.
func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
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
	forEachNode(
		doc,
		func(n *html.Node) {
			printWithDepth(depth, startElement(n))
			depth++
		},
		func(n *html.Node) {
			depth--
			printWithDepth(depth, endElement(n))
		},
	)
	return nil
}

//!-
