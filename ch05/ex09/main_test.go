package main

import "testing"

func TestExpand(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		function func(string) string
		expected string
	}{
		{
			desc:  "Simple",
			input: "Hello, $name",
			function: func(in string) string {
				return "world"
			},
			expected: "Hello, world",
		},
		{
			desc:  "Complex",
			input: "$$price",
			function: func(in string) string {
				return "10"
			},
			expected: "$10",
		},
		{
			desc:  "Multiple",
			input: "$foo $bar",
			function: func(in string) string {
				switch in {
				case "foo":
					return "FOO"
				case "bar":
					return "BAR"
				}
				return "DEFAULT"
			},
			expected: "FOO BAR",
		},
		{
			desc:  "No vars",
			input: "foo yo hello world \n",
			function: func(in string) string {
				return "NOP"
			},
			expected: "foo yo hello world \n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := Expand(tC.input, tC.function)
			if res != tC.expected {
				t.Fail()
			}
		})
	}
}
