package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	BadEcho()
	secs := time.Since(start).Seconds()
	fmt.Printf("\n\nBadEcho(): %.10fs\n\n", secs)

	start = time.Now()
	GoodEcho()
	secs = time.Since(start).Seconds()
	fmt.Printf("\n\nGoodEcho(): %.10fs\n\n", secs)
}

func BadEcho() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func GoodEcho() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
