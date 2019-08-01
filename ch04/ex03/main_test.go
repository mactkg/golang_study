package main

import "testing"

func TestReverse(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5}
	Reverse(&a)
	if a != [6]int{5, 4, 3, 2, 1, 0} {
		t.Fatalf("Error! %v", a)
	}
}
