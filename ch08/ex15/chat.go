// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const clientChannelCap int = 256

//!+broadcaster
type client struct {
	name string
	ch   chan<- string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)     // all incoming client messages
	clients  = make(map[client]bool) // all connected clients
)

func broadcaster() {
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.ch <- msg:
				default:
					// give up to send message
				}
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string, clientChannelCap) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{name: who, ch: ch}
	ch <- "You are " + who
	sb := strings.Builder{}
	sb.WriteString("Active Clients: ")
	for c := range clients {
		sb.WriteString(c.name)
		sb.WriteByte(' ')
	}
	ch <- sb.String()

	messages <- who + " has arrived"
	entering <- cli

	timeout := make(chan struct{})
	closed := make(chan struct{})

	go func() {
		duration := 5 * time.Minute
		idleTimer := time.AfterFunc(duration, func() { close(timeout) })
		input := bufio.NewScanner(conn)
		for input.Scan() {
			messages <- who + ": " + input.Text()
			idleTimer.Reset(duration)
		}
		// NOTE: ignoring potential errors from input.Err()
		idleTimer.Stop()
		close(closed)
	}()

	select {
	case <-timeout:
	case <-closed:
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
