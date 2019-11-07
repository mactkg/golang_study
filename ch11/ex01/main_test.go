package main

import (
	"bytes"
	"testing"
	"unicode/utf8"
)

type expected struct {
	counts  map[rune]int
	utflen  [utf8.UTFMax + 1]int
	invalid int
}

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		in       string
		expected expected
	}{
		{
			desc: "normal",
			in:   "hello world",
			expected: expected{
				counts: map[rune]int{
					'h': 1,
					'e': 1,
					'l': 3,
					'o': 2,
					'w': 1,
					'r': 1,
					'd': 1,
					' ': 1,
				},
				utflen: [utf8.UTFMax + 1]int{
					0, 11, 0, 0, 0,
				},
				invalid: 0,
			},
		},
		{
			desc: "にほんご",
			in:   "おはこんにちは世界〜!!",
			expected: expected{
				counts: map[rune]int{
					'お': 1,
					'は': 2,
					'こ': 1,
					'ん': 1,
					'に': 1,
					'ち': 1,
					'世': 1,
					'界': 1,
					'〜': 1,
					'!': 2,
				},
				utflen: [utf8.UTFMax + 1]int{
					0, 2, 0, 10, 0,
				},
				invalid: 0,
			},
		},
		{
			desc: "control chars",
			in:   "\t\n\t\r\n\thello",
			expected: expected{
				counts: map[rune]int{
					'\n': 2,
					'\r': 1,
					'\t': 3,
					'h':  1,
					'e':  1,
					'o':  1,
					'l':  2,
				},
				utflen: [utf8.UTFMax + 1]int{
					0, 11, 0, 0, 0,
				},
				invalid: 0,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			buf := bytes.NewBufferString(tC.in)
			counts, utflen, invalid, err := charcount(buf)
			if err != nil {
				t.Fatal(err)
			}

			if len(counts) != len(tC.expected.counts) {
				t.Fatalf("wrong length of counts: %v(expected: %v)", len(counts), len(tC.expected.counts))
			}
			for k, v := range tC.expected.counts {
				if counts[k] != v {
					t.Fatalf("wrong count of %v: %v(expected: %v)", k, counts[k], v)
				}
			}

			for k, v := range tC.expected.utflen {
				if utflen[k] != v {
					t.Fatalf("wrong count of utflen %v: %v(expected: %v)", k, utflen[k], v)
				}
			}

			if invalid != tC.expected.invalid {
				t.Fatalf("wrong invalid count: %v(expected: %v)", invalid, tC.expected.invalid)
			}
		})
	}
}
