package main

import (
	"sort"
	"testing"
)

func TestHasAndAdd(t *testing.T) {
	var s IntSet

	s.Add(10)
	if !s.Has(10) {
		t.Fatalf("%v", s.String())
	}
}

func TestString(t *testing.T) {
	var s IntSet
	s.AddAll(10, 80)

	if s.String() != "{10 80}" {
		t.Fatalf("%v", s.String())
	}
}

func TestRemove(t *testing.T) {
	var s IntSet
	s.AddAll(100, 30, 23)
	s.Remove(30)

	if s.Len() != 2 || !s.Has(100) || !s.Has(23) {
		t.Fatalf("%v", s.String())
	}
}

func TestElems(t *testing.T) {
	var s IntSet
	s.AddAll(3, 4, 5, 10, 100)

	elems := s.Elems()
	sort.Ints(elems)

	expected := []int{3, 4, 5, 10, 100}

	for i, v := range elems {
		if v != expected[i] {
			t.Fatalf("%v", s.String())
		}
	}
}
