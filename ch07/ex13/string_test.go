package eval

import (
	"fmt"
	"testing"
)

//!+Eval
func TestString(t *testing.T) {
	tests := []struct {
		expr string
	}{
		{"sqrt(A / pi)"},
		{"pow(x, 3) + pow(y, 3)"},
		{"5 / 9 * (F - 32.5)"},
		{"-1 + -x"},
		{"-1 - x"},
		{"-0.24 - x"},
	}
	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		fmt.Println(expr)

		expr2, err := Parse(expr.String())
		if err != nil {
			t.Error(err)
			continue
		}

		if fmt.Sprint(expr) != fmt.Sprint(expr2) {
			t.Errorf("%s != %s", expr, expr2)
		}
	}
}

//!-Eval

/*
//!+output
sqrt(A / pi)
	map[A:87616 pi:3.141592653589793] => 167

pow(x, 3) + pow(y, 3)
	map[x:12 y:1] => 1729
	map[x:9 y:10] => 1729

5 / 9 * (F - 32)
	map[F:-40] => -40
	map[F:32] => 0
	map[F:212] => 100
//!-output

// Additional outputs that don't appear in the book.

-1 - x
	map[x:1] => -2

-1 + -x
	map[x:1] => -2
*/
