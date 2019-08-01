package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	Reverse(&a)
	fmt.Println(a)
}

func Reverse(s *[6]int) {
	for i, j := 0, 5; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
