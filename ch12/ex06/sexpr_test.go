// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"reflect"
	"testing"
)

func TestZeroValue(t *testing.T) {
	type Tester struct {
		Integer  int
		Float    float64
		Str      string
		Boolean  bool
		IntArray []int
	}

	zero := Tester{}
	nonzero := Tester{10, 10.42, "hi", true, []int{0, 1, 2, 0, 3}}

	// zero values
	data, err := Marshal(zero)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Zero values:\n%v", string(data))

	var res Tester
	err = Unmarshal(data, &res)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, zero) {
		t.Fatal("not equal")
	}

	// non zero values
	data, err = Marshal(nonzero)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Non zero values:\n%v", string(data))

	err = Unmarshal(data, &res)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, nonzero) {
		t.Fatal("not equal")
	}
}
