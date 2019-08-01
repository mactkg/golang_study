package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts, types, utflen, invalid := CharCount(os.Stdin)

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	fmt.Printf("\ntype\tcount\n")
	for c, n := range types {
		fmt.Printf("%v\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func CharCount(input io.Reader) (map[rune]int, map[string]int, [utf8.UTFMax + 1]int, int) {
	counts := make(map[rune]int)
	types := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(input)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsControl(r):
			types["control"]++
		case unicode.IsDigit(r):
			types["digit"]++
		case unicode.IsNumber(r):
			types["number"]++
		case unicode.IsLetter(r):
			types["latter"]++
		case unicode.IsGraphic(r):
			types["graphic"]++
		default:
			types["other"]++
		}

		counts[r]++
		utflen[n]++
	}

	return counts, types, utflen, invalid
}
