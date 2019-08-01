package main

import "fmt"

func main() {
	a := []string{"abc", "abc", "ab", "ab", "def"}
	fmt.Println(a)
	fmt.Println(RemoveDupNeighbor(a))
}

func RemoveDupNeighbor(a []string) []string {
	cnt := 0
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			copy(a[i:], a[i+1:])
			cnt++
		}
	}
	return a[:len(a)-cnt]
}
