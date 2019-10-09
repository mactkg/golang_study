// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "port number")
	flag.Parse()

	addr := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(os.Stdout, addr, log.LstdFlags)
	logger.Printf("Start listening at %s", addr)

	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		logger.Printf("New connection! from %s", conn.RemoteAddr())
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
