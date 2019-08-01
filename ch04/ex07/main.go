package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := []byte("リバースthis")
	fmt.Println(string(a))
	fmt.Println(string(Reverse(a)))
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse(a []byte) []byte {
	for i := 0; i < len(a); {
		_, s := utf8.DecodeRune(a[i:])
		reverse(a[i : i+s])
		i += s
	}
	reverse(a)
	return a
}
