// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"bytes"
	"io"
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

func TestConsume(t *testing.T) {
	type Tester struct {
		Id   int
		Name string
		Attr map[string]string
	}
	data := Tester{
		Id:   88,
		Name: "hello",
		Attr: map[string]string{
			"age":   "10",
			"state": "Tokyo",
		},
	}

	// encode Test data to decode
	buf := &bytes.Buffer{}
	encoder := NewEncoder(buf)
	err := encoder.Encode(data)
	if err != nil {
		t.Fatal(err)
	}

	// pass encoded data to decoder, and restructure it
	decoder := NewDecoder(buf)
	str := ""
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}

		switch token := token.(type) {
		case String:
			str += `"` + token.String() + `"`
		case Int:
			str += token.String()
		case Symbol:
			str += token.String() + " "
		case StartList:
			str += "("
		case EndList:
			str += ")"
		}
	}

	// Pass restructured data to Unmarshal(), and compare them deeply
	tester := Tester{}
	err = Unmarshal([]byte(str), &tester)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(tester, data) {
		t.Fatalf("%v != %v", tester, data)
	}
}
