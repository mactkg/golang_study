package main

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	b := &bytes.Buffer{}
	counter, count := CountingWriter(b)
	counter.Write([]byte("Hello, world\n"))
	if *count != int64(13) {
		t.Fatal()
	}
}
