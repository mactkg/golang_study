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
	"time"
)

//!+broadcaster
type client struct {
	r <-chan string // an ingoing message channel
	w chan<- string // an outgoing message channel
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
	timeout  = 10 * time.Second// 5 * time.Minute
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.w <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			msg := "User list:\n"
			for c, _ := range clients {
				msg += fmt.Sprintf("\t%s\n", c.name)
			}
			cli.w <- msg

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.w)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	out := make(chan string, 10) // outgoing client messages
	in := make(chan string) // ingoing client messages
	go clientWriter(conn, out)
	go clientReader(conn, in)

	out<-"What's your name?: "
	who := <-in

	user := client{in, out, who}
	out <- "You are " + who
	messages <- who + " has arrived"
	entering <- user


READ:
	for {
		select {
		case str := <-in:
			messages <- who + ": " + str
		case <-time.After(timeout):
			break READ
		}
	}

	leaving <- user
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func clientReader(conn net.Conn, ch chan<- string) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ch <- scanner.Text()
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
