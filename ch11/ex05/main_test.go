package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	s, sep := "a:b:c", ":"
	words := strings.Split(s, sep)
	if got, want := len(words), 3; got != want {
		t.Errorf("Split(%q, %q) returned %d words, word %d", s, sep, got, want)
	}
}

func TestSplitTDT(t *testing.T) {
	testCases := []struct {
		desc     string
		str      string
		sep      string
		expected int
	}{
		{
			desc:     "",
			str:      "a:b:c",
			sep:      ":",
			expected: 3,
		},
		{
			desc:     "",
			str:      "a:b:c",
			sep:      "",
			expected: 5,
		},
		{
			desc:     "",
			str:      "abc",
			sep:      "",
			expected: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			words := strings.Split(tC.str, tC.sep)
			if got, want := len(words), tC.expected; got != want {
				t.Errorf("Split(%q, %q) returned %d words, word %d", tC.str, tC.sep, got, want)
			}
		})
	}
}
