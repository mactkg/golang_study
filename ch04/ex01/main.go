package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func main() {
	a, b := sha256.Sum256([]byte("24")), sha256.Sum256([]byte("42"))
	fmt.Printf("%x\n%x\n%d\n", a, b, a, b, Sha256HashDiff(a, b))
}

// Comes from ch02/ex05
func popCountLoopWithOutShift(x uint64) int {
	var c int

	for x != 0 {
		x = x & (x - 1)
		c++
	}
	return c
}

func Sha256HashDiff(a, b [32]byte) int {
	sum := 0
	for i := 0; i < 4; i++ {
		c := a[8*i : 8*(i+1)]
		d := b[8*i : 8*(i+1)]

		x := binary.LittleEndian.Uint64(c)
		y := binary.LittleEndian.Uint64(d)

		sum += popCountLoopWithOutShift(x ^ y)
	}
	return sum
}
