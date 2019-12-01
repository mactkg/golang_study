// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"reflect"
	"regexp"
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

func TestPrettyPrint(t *testing.T) {
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
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Pretty-print it:
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	role := `(Dr. Strangelove|Grp\. Capt\. Lionel Mandrake|Pres\. Merkin Muffley|Gen\. Buck Turgidson|Brig\. Gen\. Jack D\. Ripper|Maj\. T\.J\. "King" Kong)`
	actor := `(Peter Sellers|George C\. Scott|Sterling Hayden|Slim Pickens)`
	expect := `\(\(Title "Dr. Strangelove"\)
 \(Subtitle "How I Learned to Stop Worrying and Love the Bomb"\)
 \(Year 1964\)
 \(Actor \(\("` + role + `" "` + actor + `"\)
         \("` + role + `" "` + actor + `"\)
         \("` + role + `" "` + actor + `"\)
         \("` + role + `" "` + actor + `"\)
         \("` + role + `" "` + actor + `"\)\)\)
 \(Oscars \("Best Actor \(Nomin\.\)"
          "Best Adapted Screenplay \(Nomin\.\)"
          "Best Director \(Nomin\.\)"
          "Best Picture \(Nomin\.\)"\)\)
 \(Sequel nil\)\)`
	r := regexp.MustCompile(expect)
	if !r.MatchString(string(data)) {
		t.Fatalf("Wrong indent.\nGot:\n%s\n\nExpected:\n%s\n", data, expect)
	}
}

func TestPPArray(t *testing.T) {
	data, err := Marshal(struct {
		ia []int
	}{ia: []int{1, 2, 3, 4}})
	if err != nil {
		t.Fatal(err)
	}
	expect := `((ia (1
      2
      3
      4)))`
	if string(data) != expect {
		t.Fatalf("Wrong indent.\nGot:\n%s\n\nExpected:\n%s\n", data, expect)
	}
}

func TestPPStruct(t *testing.T) {
	data, err := Marshal(struct {
		ia    int
		st    string
		strct *struct{}
	}{ia: 1, st: "string here", strct: nil})
	if err != nil {
		t.Fatal(err)
	}
	expect := `((ia 1)
 (st "string here")
 (strct nil))`
	if string(data) != expect {
		t.Fatalf("Wrong indent.\nGot:\n%s\n\nExpected:\n%s\n", data, expect)
	}
}
