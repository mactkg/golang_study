package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	a = Rotate(a, 2)
	fmt.Println(a)
}

func Rotate(a []int, l int) []int {
	buf := make([]int, l)
	copy(buf[0:], a[:l])

	left := len(a) - l
	for i := 0; i < len(a); i++ {
		if i < left {
			a[i] = a[i+l]
		} else {
			a[i] = buf[i-left]
		}
	}

	return a
}
