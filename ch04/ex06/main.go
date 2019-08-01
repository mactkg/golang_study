package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	a := []byte("Compress  Spaces")
	fmt.Println(a)
	fmt.Println(CompressSpaces(a))
}

func CompressSpaces(a []byte) []byte {
	cnt := 0
	for i := 0; i < len(a); {
		r1, s1 := utf8.DecodeRune(a[i:])
		r2, s2 := utf8.DecodeRune(a[i+s1:])

		if unicode.IsSpace(r1) && unicode.IsSpace(r2) {
			a[i] = ' '
			copy(a[i+1:], a[i+s1+s2:])
			cnt += s1 + s2
			i++
		} else {
			i += s1
		}
	}

	return a[:len(a)-cnt]
}
