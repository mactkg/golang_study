// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"reflect"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestEx12_3(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	testCases := []struct {
		desc   string
		input  interface{}
		expect string
	}{
		{
			desc: "Boolean",
			input: struct {
				t bool
				f bool
			}{t: true, f: false},
			expect: "((t t) (f nil))",
		},
		{
			desc: "Float",
			input: struct {
				v64 float64
				v32 float32
			}{v64: float64(4.2112341234), v32: float32(-4.242424242424242)},
			expect: "((v64 4.211234) (v32 -4.242424))",
		},
		{
			desc: "Complex",
			input: struct {
				comp128 complex128
				comp64  complex64
			}{comp128: complex(1, 2), comp64: complex64(complex(-1, 2))},
			expect: "((comp128 #C(1.0, 2.0)) (comp64 #C(-1.0, 2.0)))",
		},
		{
			desc: "nil",
			input: struct {
				data *Movie
			}{data: nil},
			expect: "((data nil))",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			data, err := Marshal(tC.input)
			if err != nil {
				t.Fatalf("Error: %v", err)
			}
			if string(data) != tC.expect {
				t.Fatalf("Unexpected result: %v(expected: %v)", string(data), tC.expect)
			}
		})
	}
}

func TestEx12_3_Interface(t *testing.T) {
	type Input struct {
		data interface{}
	}
	testCases := []struct {
		desc   string
		input  Input
		expect string
	}{
		{
			desc: "Interface",
			input: Input{data: struct {
				s string
				i int
			}{"foo", 42}},
			expect: "((data (\"struct { s string; i int }\" ((s \"foo\") (i 42)))))",
		}, {
			desc:   "Interface2",
			input:  Input{data: []int{0, 1, 2, 3}},
			expect: "((data (\"[]int\" (0 1 2 3))))",
		},
		{
			desc:   "Interface(nil)",
			input:  Input{data: nil},
			expect: "((data nil))",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			data, err := Marshal(tC.input)
			if err != nil {
				t.Fatalf("Error: %v", err)
			}
			if string(data) != tC.expect {
				t.Fatalf("Unexpected result: %v(expected: %v)", string(data), tC.expect)
			}
		})
	}
}
