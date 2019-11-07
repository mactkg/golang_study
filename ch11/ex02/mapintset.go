package main

import (
	"bytes"
	"fmt"
	"sort"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type MapIntSet map[uint]struct{}

// Has reports whether the set contains the non-negative value x.
func (s MapIntSet) Has(x int) bool {
	if x < 0 {
		return false
	}
	_, ok := s[uint(x)]
	return ok
}

// Add adds the non-negative value x to the set.
func (s *MapIntSet) Add(x int) {
	if x < 0 {
		return
	}
	(*s)[uint(x)] = struct{}{}
}

// UnionWith sets s to the union of s and t.
func (s *MapIntSet) UnionWith(t *MapIntSet) {
	for k, _ := range *t {
		(*s)[k] = struct{}{}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *MapIntSet) String() string {
	set := []uint{}
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, v := range *s {
		if v != struct{}{} {
			continue
		}
		set = append(set, k)
	}

	sort.Slice(set, func(i, j int) bool {
		return set[i] < set[j]
	})

	for _, v := range set {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte('}')
	return buf.String()
}
