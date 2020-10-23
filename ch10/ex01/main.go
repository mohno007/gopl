// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 287.

//!+main

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

var formatFlag = flag.String("f", "", "export format")

func main() {
	flag.Parse()

	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	switch *formatFlag {
	case "png":
		if err := png.Encode(os.Stdout, img); err != nil {
			fmt.Fprintf(os.Stderr, "png: %v\n", err)
			os.Exit(1)
		}
	case "jpeg":
		if err := jpeg.Encode(os.Stdout, img, &jpeg.Options{Quality: 95}); err != nil {
			fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
			os.Exit(1)
		}
	case "gif":
		if err := gif.Encode(os.Stdout, img, &gif.Options{NumColors: 256}); err != nil {
			fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown format: '%v'\n", err)
		os.Exit(1)
	}
}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
