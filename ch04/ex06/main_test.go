package main

import (
	"testing"
)

func TestCompressSpaces(t *testing.T) {
	a := []byte("ハロー　　Helloワールド")
	expected := []byte("ハロー Helloワールド")
	a = CompressSpaces(a)

	for i := range a {
		if a[i] != expected[i] {
			t.Fatalf("Error! Got: %b, Expected: %b", a[i], expected[i])
		}
	}
}
