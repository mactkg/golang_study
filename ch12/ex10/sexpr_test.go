// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"bytes"
	"reflect"
	"testing"
)

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
	buf := &bytes.Buffer{}
	encoder := NewEncoder(buf)
	encoder.Encode(strangelove)
	t.Logf("Marshal() = %s\n", buf)

	decoder := NewDecoder(buf)
	var movie Movie
	if err := decoder.Decode(&movie); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	t.Logf("Decode() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}

func TestEx12_10(t *testing.T) {
	type Test struct {
		F32 float32
		F64 float64
		T   bool
		F   bool
		I   interface{}
	}
	d := Test{
		F32: 123.456, // I want to test negative value, but I can't impl that
		F64: 123.456,
		T:   true,
		F:   false,
		I:   []int{1, 2, 4, 8},
	}

	RegisterInterface("[]int", reflect.TypeOf([]int{}))

	res, err := Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))

	out := Test{}
	err = Unmarshal(res, &out)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(d, out) {
		t.Fatalf("%v != %v", d, out)
	}
}
