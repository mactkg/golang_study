package main

import (
	"fmt"

	"github.com/mactkg/golang_study/ch02/ex01/tempconv"
)

func main() {
	k := tempconv.Kelvin(10)
	fmt.Println(k, tempconv.KToC(k), tempconv.KToF(k))
}
