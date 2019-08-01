package main

import "testing"

func TestSha256HashDiff(t *testing.T) {
	a := [32]byte{1, 1, 0, 2}
	b := [32]byte{}
	result := Sha256HashDiff(a, b)
	if result != 3 {
		t.Fatalf("Error! %d", result)
	}
}
