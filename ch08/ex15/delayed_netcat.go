// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

//!+
func main() {
	var d int
	flag.IntVar(&d, "delayMs", 100, "delay of read i/o in Millisecond")
	flag.Parse()
	delay := time.Millisecond * time.Duration(d)

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	tcpConn := conn.(*net.TCPConn)
	done := make(chan struct{})
	go func() {
		delayedCopy(delay, os.Stdout, conn) // NOTE: ignoring errors
		tcpConn.CloseRead()
		log.Println("// Connection to read was closed too")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	tcpConn.CloseWrite()
	log.Println("// Connection to write was closed")
	<-done // wait for background goroutine to finish
}

//!-

func delayedCopy(delay time.Duration, dst io.Writer, src io.Reader) {
	r := bufio.NewReader(src)
	w := bufio.NewWriter(dst)

	for {
		byte, err := r.ReadByte()
		if err != nil {
			fmt.Println(err)
			break
		}

		w.WriteByte(byte)
		w.Flush()

		// wait random
		<-time.After(delay)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
