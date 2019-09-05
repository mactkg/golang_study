package main

import "fmt"

func main() {
	fmt.Println(returnNonZeroValueWithoutReturn())
}

func returnNonZeroValueWithoutReturn() (res int) {
	defer func() {
		recover()
		res = 100
	}()
	panic(1)
}
