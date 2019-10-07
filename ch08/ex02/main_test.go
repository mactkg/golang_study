package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
)

func sendAndRead(c net.Conn, str string, code string) (string, error) {
	scanner := bufio.NewScanner(c)
	_, err := c.Write([]byte(str + "\n"))
	if err != nil {
		return "", err
	}

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

func TestGoldenPath(t *testing.T) {
	reciever, sender := net.Pipe()
	scanner := bufio.NewScanner(sender)
	go handleConn(reciever)

	ok := scanner.Scan()
	text := scanner.Text()
	if !ok || strings.Index(text, "220") != 0 {
		t.Fatalf("Connection Failed: %v", text)
	}

	// login
	_, err := sendAndRead(sender, "USER kenta", "230")
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

	// change structure to file
	_, err = sendAndRead(sender, "STRU F", "200")
	if err != nil {
		t.Fatalf("Failed changing structure! (%v)", err)
	}

	_, err = sendAndRead(sender, "QUIT", "221")
	if err != nil {
		t.Fatalf("Quit failed! (%v)", err)
	}

	sender.Close()
}
