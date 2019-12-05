package main

import "testing"

func TestCheckCycled(t *testing.T) {
	type link struct {
		value interface{}
		next  *link
	}

	var a, b, c link
	a = link{"A", &b}
	b = link{10, &c}
	c = link{"c", &a}

	var d []interface{}
	d = []interface{}{"a", "b", &d}

	var e map[string]interface{}
	e = map[string]interface{}{
		"foo": a,
	}

	testCases := []struct {
		desc     string
		value    interface{}
		expected bool
	}{
		{
			desc:     "atoms",
			value:    []interface{}{0, 1, "foo", 0.4, false},
			expected: false,
		},
		{
			desc: "noloop",
			value: link{
				value: "foo",
				next: &link{
					value: "bar",
					next:  nil,
				},
			},
			expected: false,
		},
		{
			desc:     "loop",
			value:    a,
			expected: true,
		},
		{
			desc:     "loop interface",
			value:    d,
			expected: true,
		},
		{
			desc:     "loop map",
			value:    e,
			expected: true,
		},
		{
			desc:     "no loop array",
			value:    [4]int{0, 2, 4, 8},
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if CheckCycled(tC.value) != tC.expected {
				t.Fatalf("error: %v(%s)", tC.value, tC.desc)
			}
		})
	}
}
