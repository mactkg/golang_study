package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	i := IntSet{}
	m := MapIntSet{}

	i.Add(1)
	i.Add(2)
	i.Add(3)
	i.Add(math.MaxUint32)
	i.Add(0)

	m.Add(1)
	m.Add(2)
	m.Add(3)
	m.Add(math.MaxUint32)
	m.Add(0)

	if i.String() != m.String() {
		t.Fatalf("Error! IntSet: %v, MapIntSet: %v", i.String(), m.String())
	}
}

func TestHas(t *testing.T) {
	i := IntSet{}
	m := MapIntSet{}

	i.Add(1)
	i.Add(3)
	i.Add(2)
	i.Add(math.MaxUint32)

	m.Add(1)
	m.Add(3)
	m.Add(2)
	m.Add(math.MaxUint32)

	if i.Has(1) != m.Has(1) {
		t.Fatalf("Error! IntSet.Has(1): %v, MapIntSet.Has(1): %v", i.Has(1), m.Has(1))
	}

	if i.Has(math.MaxUint32) != m.Has(math.MaxUint32) {
		t.Fatalf("Error! IntSet.Has(MaxUint32): %v, MapIntSet.Has(MaxUint32): %v", i.Has(math.MaxUint32), m.Has(math.MaxUint32))
	}

	// panic
	// 	if i.Has(-100) != m.Has(-100) {

	// 	}
}

func TestUnionWith(t *testing.T) {
	i1 := IntSet{}
	i2 := IntSet{}
	m1 := MapIntSet{}
	m2 := MapIntSet{}

	i1.Add(1)
	i1.Add(3)
	i2.Add(2)
	i2.Add(3)
	i2.Add(math.MaxUint32)
	i1.UnionWith(&i2)

	m1.Add(1)
	m1.Add(3)
	m2.Add(2)
	m2.Add(3)
	m2.Add(math.MaxUint32)
	m1.UnionWith(&m2)

	if i1.String() != m1.String() {
		t.Fatalf("Error! IntSet: %v, MapIntSet: %v", i1, m1)
	}
}
