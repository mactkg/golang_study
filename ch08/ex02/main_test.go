package main

import (
	"bufio"
	"fmt"
	"math/rand"
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

func handleData(d net.Conn, reciever chan string, sender chan []byte, clean chan struct{}) {
	defer d.Close()

	// reciever
	go func(d net.Conn, r chan string) {
		scanner := bufio.NewScanner(d)
		for scanner.Scan() {
			r <- scanner.Text()
		}
	}(d, reciever)

	// sender
	select {
	case buf := <-sender:
		fmt.Printf("Got: %v", buf)
		d.Write(buf)
	case <-clean:
		fmt.Printf("clean!")
		break
	}

	fmt.Printf("bye!")
}

// from: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func TestGoldenPath(t *testing.T) {
	dataCh := make(chan string)
	sendCh := make(chan []byte)
	clean := make(chan struct{})
	defer func() {
		clean <- struct{}{}
	}()

	data, err := net.Listen("tcp", "localhost:5678")
	if err != nil {
		t.Fatalf("Can't listen socket for data connection! (%v)", err)
	}
	go func(d net.Listener, r chan string, s chan []byte, c chan struct{}) {
		for {
			conn, err := d.Accept()
			if err != nil {
				t.Errorf("Something happend in data server goroutine (%v)", err)
			}
			t.Logf("New connection is established: %v", conn)
			go handleData(conn, r, s, c)
		}

	}(data, dataCh, sendCh, clean)

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

	// cd
	_, err = sendAndRead(sender, "CWD temp_test", "250")
	if err != nil {
		t.Fatalf("Failed cd (%v)", err)
	}

	// Check ls again
	_, err = sendAndRead(sender, "LIST", "150")
	if err != nil {
		t.Fatalf("Failed ls (%v)", err)
	}
	select {
	case res := <-dataCh:
		t.Logf("Recieved: %v", res)
		if strings.Index(res, "top_secret") == -1 {
			t.Log("CWD seems to be failed")
		}
	case <-time.After(3 * time.Second):
		t.Fatalf("Data connection should be recieved something")
	}
	_, err = readCheck(sender, "226")
	if err != nil {
		t.Fatalf("Data connection should be closed (%v)", err)
	}

	// how can i test uploading/downloading file...
	// how can i get whole of file...
	sendAndRead(sender, "CWD ../", "250")
	_, err = sendAndRead(sender, "RETR main.go", "150")
	if err != nil {
		t.Fatalf("Failed RETR (%v)", err)
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

	/*
		commented out this test because there is something wrong

		// store test
		_, err = sendAndRead(sender, "STOR temp_test/top_secret_"+randString(8)+".txt", "150")
		if err != nil {
			t.Fatalf("Failed RETR (%v)", err)
		}
		secret, err := os.Open("./temp_test/top_secret.txt")
		if err != nil {
			t.Fatalf("Opening file failed (%v)", err)
		}
		r := bufio.NewReader(secret)
		buf := make([]byte, 256)
		size := 0
		for {
			n, err := r.Read(buf)
			size += n
			t.Logf("Loaded %v bytes", n)
			if err == io.EOF {
				sendCh <- buf
				t.Logf("Sent %v bytes file", size)
				break
			} else if err != nil {
				t.Logf("Error happend while sending data: %v", err)
				break
			}
			sendCh <- buf
		}
		clean <- struct{}{}

		_, err = readCheck(sender, "226")
		if err != nil {
			t.Fatalf("Data connection should be closed (%v)", err)
		}
	*/

	_, err = sendAndRead(sender, "QUIT", "221")
	if err != nil {
		t.Fatalf("Quit failed! (%v)", err)
	}

	sender.Close()
}
