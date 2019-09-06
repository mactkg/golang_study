package main

import (
	"testing"
)

func TestWordCounter(t *testing.T) {
	testCases := []struct {
		desc     string
		data     string
		expected int
	}{
		{
			desc:     "Simple",
			data:     "Hello world yo-yo.",
			expected: 3,
		},
		{
			desc:     "Multiple Lines",
			data:     "Hello world,\nI'm from Japan.",
			expected: 5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var c WordCounter
			c.Write([]byte(tC.data))
			if c != WordCounter(tC.expected) {
				t.Fatalf("Expected: %v, Got: %v", tC.expected, c)
			}
		})
	}
}

func TestLineCounter(t *testing.T) {
	testCases := []struct {
		desc     string
		data     string
		expected int
	}{
		{
			desc:     "Simple",
			data:     "Hello world yo-yo.",
			expected: 1,
		},
		{
			desc:     "Multiple Lines",
			data:     "Hello world,\nI'm from Japan.\nGood\n\nMorning!!",
			expected: 5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var c LineCounter
			c.Write([]byte(tC.data))
			if c != LineCounter(tC.expected) {
				t.Fatalf("Expected: %v, Got: %v", tC.expected, c)
			}
		})
	}
}
