package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func clock(name string, c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(c)
	for {
		str, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Printf("[%s] %s\n", name, str)
	}

	fmt.Println("bye")
}

func main() {
	type remote struct {
		name string
		addr string
	}
	remotes := []remote{}

	for _, addr := range os.Args[1:] {
		res := strings.Split(addr, "=")
		remotes = append(remotes, remote{name: res[0], addr: res[1]})
	}
	fmt.Println(remotes)

	for _, r := range remotes {
		conn, err := net.Dial("tcp", r.addr)
		if err != nil {
			fmt.Printf("Connection error: %v\n", err)
			continue
		}

		go clock(r.name, conn)
	}

	time.Sleep(time.Minute)
}
