package main

import (
	"sort"
	"testing"
)

func TestElems(t *testing.T) {
	var s IntSet
	s.AddAll(3, 4, 5, 10, 100)

	elems := s.Elems()
	sort.Ints(elems)

	expected := []int{3, 4, 5, 10, 100}

	for i, v := range elems {
		if v != expected[i] {
			t.Fatal()
		}
	}
}
