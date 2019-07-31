package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// better comma
func main() {
	fmt.Println(Comma(os.Args[1]))
}

// Comma add comma to number and return it as string
func Comma(s string) string {
	dot := strings.Index(s, ".")
	n := len(s)

	var buf bytes.Buffer
	if dot < 0 {
		// integer
		n := len(s)
		if n <= 3 {
			return s
		}

		shift := (n - 1) % 3
		for i := 0; i < n; i++ {
			buf.WriteByte(s[i])
			if i%3-shift == 0 && i < n-1 {
				buf.WriteString(",")
			}
		}
	} else {
		// float
		if dot <= 3 {
			return s
		}

		shift := (dot - 1) % 3
		for i := 0; i < n; i++ {
			buf.WriteByte(s[i])
			if i%3-shift == 0 && i < dot-1 {
				buf.WriteString(",")
			}
		}
	}

	return buf.String()
}
