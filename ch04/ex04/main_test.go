package main

import "testing"

func TestRotate(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5}
	expected := []int{3, 4, 5, 0, 1, 2}
	a = Rotate(a, 3)

	for i := range a {
		if a[i] != expected[i] {
			t.Fatalf("Error! Got: %d, Expected: %d", a[i], expected[i])
		}
	}
}
