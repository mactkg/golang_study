package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMyLimitReader(t *testing.T) {
	data := strings.NewReader("abcdefg")
	limitReader := MyLimitReader(data, 3)
	buf := make([]byte, 3)
	n, _ := limitReader.Read(buf)

	// abc
	if n != 3 {
		t.Fatal()
	}

	if bytes.Compare(buf, []byte("abc")) != 0 {
		t.Fatalf("Expected: %v, Got: %v", []byte("abc"), buf)
	}

	// def
	buf = make([]byte, 3)
	n, _ = limitReader.Read(buf)

	if n != 3 {
		t.Fatal()
	}

	if bytes.Compare(buf, []byte("def")) != 0 {
		t.Fatalf("Expected: %v, Got: %v", []byte("def"), buf)
	}

	// g
	buf = make([]byte, 3)
	n, _ = limitReader.Read(buf)

	if n != 1 {
		t.Fatal()
	}

	if bytes.Compare(buf, []byte{103, 0, 0}) != 0 { // [g, `empty`, `empty`]
		t.Fatalf("Expected: %v, Got: %v", []byte{103, 0, 0}, buf)
	}
}
