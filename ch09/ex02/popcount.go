package popcount

import "sync"

var genPCOnce sync.Once
var pc [256]byte
var ppc [256]byte

func init() {
	for i := range pc {
		ppc[i] = ppc[i/2] + byte(i&1)
	}
}

func genPC() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountSingle(x uint64) int {
	return int(ppc[byte(x>>(0*8))] +
		ppc[byte(x>>(1*8))] +
		ppc[byte(x>>(2*8))] +
		ppc[byte(x>>(3*8))] +
		ppc[byte(x>>(4*8))] +
		ppc[byte(x>>(5*8))] +
		ppc[byte(x>>(6*8))] +
		ppc[byte(x>>(7*8))])
}

func PopCountSingleLazy(x uint64) int {
	genPCOnce.Do(genPC)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var res int
	var i uint

	for ; i < 8; i++ {
		res = res + int(pc[byte(x>>(i*8))])
	}
	return res
}

func PopCountLoopShift(x uint64) int {
	var res int
	var i uint

	for ; i < 64; i++ {
		res = res + int(x&1)
		x = x >> 1
	}
	return res
}

func PopCountLoopWithOutShift(x uint64) int {
	var c int

	for x != 0 {
		x = x & (x - 1)
		c++
	}
	return c
}
