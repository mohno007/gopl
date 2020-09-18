// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c *net.TCPConn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	// NOTE: ignoring potential errors from input.Err()
	for input.Scan() {
		text := input.Text()
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			echo(c, s, 1*time.Second)
		}(text)
	}
	go func() {
		wg.Wait()
		c.Close()
	}()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		tcpConn, _ := conn.(*net.TCPConn)
		go handleConn(tcpConn)
	}
}
