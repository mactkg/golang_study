package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileStore := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileStore)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileStore)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, fileStore[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileStore map[string][]string) {
	input := bufio.NewScanner(f)
	starring := make(map[string]bool)
	for input.Scan() {
		counts[input.Text()]++
		starring[input.Text()] = true
	}
	for k := range starring {
		fileStore[k] = append(fileStore[k], f.Name())
	}
}
