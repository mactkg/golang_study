package main

import (
	"testing"
)

func TestCompressSpaces(t *testing.T) {
	a := []byte("ハローワールド")
	expected := []byte("ドルーワーロハ")
	a = Reverse(a)

	for i := range a {
		if a[i] != expected[i] {
			t.Fatalf("Error! Got: %b, Expected: %b", a[i], expected[i])
		}
	}
}
