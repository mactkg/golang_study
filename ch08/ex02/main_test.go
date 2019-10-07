package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

func readCheck(c net.Conn, code string) (string, error) {
	scanner := bufio.NewScanner(c)
	ok := scanner.Scan()
	if !ok {
		return "", fmt.Errorf("Can't recieve data")
	}

	text := scanner.Text()
	if strings.Index(text, code) != 0 {
		return text, fmt.Errorf("Wrong code: %v", text)
	}
	return text, nil
}

func sendAndRead(c net.Conn, str string, code string) (string, error) {
	_, err := c.Write([]byte(str + "\n"))
	if err != nil {
		return "", err
	}

	return readCheck(c, code)
}

func handleData(d net.Conn, sender chan string) {
	defer d.Close()

	scanner := bufio.NewScanner(d)
	for scanner.Scan() {
		sender <- scanner.Text()
	}
}

func TestGoldenPath(t *testing.T) {
	dataCh := make(chan string)
	data, err := net.Listen("tcp", "localhost:5678")
	if err != nil {
		t.Fatalf("Can't listen socket for data connection! (%v)", err)
	}
	go func(d net.Listener, ch chan string) {
		for {
			conn, err := d.Accept()
			if err != nil {
				t.Errorf("Something happend in data server goroutine (%v)", err)
			}
			t.Logf("New connection is established: %v", conn)
			go handleData(conn, ch)
		}

	}(data, dataCh)

	reciever, sender := net.Pipe()
	go handleConn(reciever)

	// welcome
	_, err = readCheck(sender, "220")
	if err != nil {
		t.Fatalf("FTP Server should be return welcome message")
	}

	// login
	_, err = sendAndRead(sender, "USER kenta", "230")
	if err != nil {
		t.Fatalf("Login failed! (%v)", err)
	}

	// change type to ascii, non print
	_, err = sendAndRead(sender, "TYPE A N", "200")
	if err != nil {
		t.Fatalf("Failed changing type! (%v)", err)
	}

	// change mode to stream
	_, err = sendAndRead(sender, "MODE S", "200")
	if err != nil {
		t.Fatalf("Failed changing mode! (%v)", err)
	}

	// NOOP!
	_, err = sendAndRead(sender, "NOOP", "200")
	if err != nil {
		t.Fatalf("NOOP command should be return 200 (%v)", err)
	}

	// change structure to file
	_, err = sendAndRead(sender, "STRU F", "200")
	if err != nil {
		t.Fatalf("Failed changing structure! (%v)", err)
	}

	// change port
	_, err = sendAndRead(sender, "PORT 127,0,0,1,22,46", "200")
	if err != nil {
		t.Fatalf("Failed changing addr for port! (%v)", err)
	}

	// Check ls
	_, err = sendAndRead(sender, "LIST", "150")
	if err != nil {
		t.Fatalf("Failed ls (%v)", err)
	}
	select {
	case res := <-dataCh:
		t.Logf("Recieved: %v", res)
	case <-time.After(3 * time.Second):
		t.Fatalf("Data connection should be recieved something")
	}
	_, err = readCheck(sender, "226")
	if err != nil {
		t.Fatalf("Data connection should be closed (%v)", err)
	}

	// Check ls again
	_, err = sendAndRead(sender, "LIST /usr/local/", "150")
	if err != nil {
		t.Fatalf("Failed ls (%v)", err)
	}
	select {
	case res := <-dataCh:
		t.Logf("Recieved: %v", res)
	case <-time.After(3 * time.Second):
		t.Fatalf("Data connection should be recieved something")
	}
	_, err = readCheck(sender, "226")
	if err != nil {
		t.Fatalf("Data connection should be closed (%v)", err)
	}

	_, err = sendAndRead(sender, "QUIT", "221")
	if err != nil {
		t.Fatalf("Quit failed! (%v)", err)
	}

	sender.Close()
}
