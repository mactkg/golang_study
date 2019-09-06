package main

import (
	"bytes"
	"testing"
)

func TestMyReader(t *testing.T) {
	data := "abcdefg\n123456"
	expected := []byte("abcdefg\n123456")

	reader := MyReader(data)
	buf := make([]byte, 1024)
	len, _ := reader.Read(buf)

	if bytes.Compare(buf[:len], expected) != 0 {
		t.Fatalf("Expected: %v, Got: %v", expected, buf)
	}
}
