// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	timezones := make([]string, 0, len(os.Args)-1)
	scanners := make(map[string]bufio.Scanner)

	for _, arg := range os.Args[1:] {
		result := strings.Split(arg, "=")
		if len(result) != 2 {
			log.Fatalf("clock format is invalid %s (expected format: TZ=ADDRESS:PORT)", arg)
		}

		timezone := result[0]
		address := result[1]
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		scanner := bufio.NewScanner(conn)
		scanner.Split(bufio.ScanLines)

		timezones = append(timezones, timezone)
		scanners[timezone] = *scanner
	}

	for {
		for _, timezone := range timezones {
			scanner := scanners[timezone]
			ok := scanner.Scan()
			fmt.Printf("\033[K")
			if ok {
				fmt.Printf("%s\t\t%s\n", timezone, scanner.Text())
			} else {
				fmt.Printf("%s\t\tNA (reason: %v)\n", timezone, scanner.Err())
			}
		}
		fmt.Printf("\033[%dA", len(timezones))
	}
}

//!-
