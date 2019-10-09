package main

import (
	"sort"
	"testing"
)

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		result bool
	}{
		{
			desc:   "ok",
			input:  []int{1, 2, 3, 2, 1},
			result: true,
		},
		{
			desc:   "ok2",
			input:  []int{0},
			result: true,
		},
		{
			desc:   "bad",
			input:  []int{1, 1, 1, 3, 1},
			result: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if isPalindrome(sort.IntSlice(tC.input)) != tC.result {
				t.Fail()
			}
		})
	}
}
