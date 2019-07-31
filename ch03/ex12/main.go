package main

import (
	"fmt"
	"os"
	"sort"
)

// better comma
func main() {
	fmt.Println(CheckAnagram(os.Args[1], os.Args[2]))
}

// CheckAnagram checks two strings are anagram and return
func CheckAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	r1, r2 := []rune(s1), []rune(s2)
	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})
	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}
