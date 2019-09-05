package main

import "testing"

func TestTopoSort(t *testing.T) {
	testCases := []struct {
		desc        string
		shouldError bool
		data        map[string][]string
	}{
		{
			desc:        "normal",
			shouldError: false,
			data: map[string][]string{
				"A": {"B"},
				"B": {"C"},
			},
		},
		{
			desc:        "circular references",
			shouldError: true,
			data: map[string][]string{
				"A": {"B"},
				"B": {"A"},
				"C": {"D"},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := topoSort(tC.data)
			if (err != nil) != tC.shouldError {
				t.Fatalf("toposort should raise error when this test case: %v", tC.data)
			}
		})
	}
}
