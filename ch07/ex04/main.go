package main

import (
	"fmt"
	"io"
)

func main() {
	r := MyReader("abcdefg")
	buf := make([]byte, 1024)
	r.Read(buf)
	fmt.Printf("%s\n", string(buf))
}

type StringReader string

func (r StringReader) Read(p []byte) (int, error) {
	copy(p, r)
	return len(r), nil
}

func MyReader(str string) io.Reader {
	r := StringReader(str)
	return r
}
