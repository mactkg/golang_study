package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	counts := Wordfreq(os.Stdin)

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

func Wordfreq(input io.Reader) map[string]int {
	counts := make(map[string]int)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "wordfreq	: %v\n", err)
			os.Exit(1)
		}

		w := scanner.Text()
		counts[strings.Trim(w, ".!?")]++
	}

	return counts
}
