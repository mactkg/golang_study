package main

import (
	"bufio"
	"bytes"
)

func main() {

}

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, err
		}

		*c += WordCounter(1)
	}
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, err
		}

		*c += LineCounter(1)
	}
	return len(p), nil
}
