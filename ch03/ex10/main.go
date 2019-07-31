package main

import (
	"bytes"
	"fmt"
	"os"
)

// better comma
func main() {
	fmt.Println(Comma(os.Args[1]))
}

// Comma add comma to number and return it as string
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	shift := (n - 1) % 3
	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])
		if i%3-shift == 0 && i < n-1 {
			buf.WriteString(",")
		}
	}

	return buf.String()
}
