package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	response := make(chan string)
	cancel := make(chan struct{})
	send := func(hostname string) {
		res, err := request(hostname, cancel)
		if err != nil {
			fmt.Println(err)
			return
		}
		response<-res
	}

	go send("https://httpbin.org/delay/4")
	go send("https://httpbin.org/delay/1?dummy")
	go send("https://httpbin.org/delay/1")

	fmt.Println(<-response)
	close(cancel)
}
func request(url string, cancel <-chan struct{}) (response string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Cancel = cancel // deprecated, but this exercise is for channel

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	io.Copy(buf, res.Body)
	return buf.String(), nil
}
