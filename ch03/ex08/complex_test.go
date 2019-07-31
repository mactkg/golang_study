package main

import (
	"math/big"
	"testing"
)

// Big.Float

func TestMultiplyBigFloatComplex(t *testing.T) {
	// (0 + i) * (1 + i) == (-1 + i)
	var (
		aReal = big.NewFloat(0)
		aImag = big.NewFloat(1)
		bReal = big.NewFloat(1)
		bImag = big.NewFloat(1)
	)

	resultReal, resultImag := MultiplyBigFloatComplex(aReal, aImag, bReal, bImag)
	if resultReal.Cmp(big.NewFloat(-1)) != 0 || resultImag.Cmp(big.NewFloat(1)) != 0 {
		t.Fatalf("Error. (%v + %vi)", resultReal, resultImag)
	}
}

func TestMultiplyBigFloatComplexWithSameNumbers(t *testing.T) {
	// (0 + i) * (0 + i) == (-1)
	var (
		aReal = big.NewFloat(0)
		aImag = big.NewFloat(1)
	)

	resultReal, resultImag := MultiplyBigFloatComplex(aReal, aImag, aReal, aImag)
	if resultReal.Cmp(big.NewFloat(-1)) != 0 || resultImag.Cmp(big.NewFloat(0)) != 0 {
		t.Fatalf("Error. (%v + %vi)", resultReal, resultImag)
	}
}

func TestSquareBigFloatComplexWithSameNumbers(t *testing.T) {
	// (0 + i) * (0 + i) == (-1)
	var (
		aReal = big.NewFloat(0)
		aImag = big.NewFloat(1)
	)

	resultReal, resultImag := SquareBigFloatComplex(aReal, aImag)
	if resultReal.Cmp(big.NewFloat(-1)) != 0 || resultImag.Cmp(big.NewFloat(0)) != 0 {
		t.Fatalf("Error. (%v + %vi)", resultReal, resultImag)
	}
}

func TestAbsBigFloatComplex(t *testing.T) {
	// |(3 + 4i)| = 5
	var (
		aReal = big.NewFloat(3)
		aImag = big.NewFloat(4)
	)

	result := AbsBigFloatComplex(aReal, aImag)
	if result.Cmp(big.NewFloat(5)) != 0 {
		t.Fatalf("Error. %v", result)
	}
}

// Big.Rat

func TestMultiplyBigRatComplex(t *testing.T) {
	// (0 + i) * (1 + i) == (-1 + i)
	var (
		aReal = big.NewRat(0, 1)
		aImag = big.NewRat(1, 1)
		bReal = big.NewRat(1, 1)
		bImag = big.NewRat(1, 1)
	)

	resultReal, resultImag := MultiplyBigRatComplex(aReal, aImag, bReal, bImag)
	if resultReal.Cmp(big.NewRat(-1, 1)) != 0 || resultImag.Cmp(big.NewRat(1, 1)) != 0 {
		t.Fatalf("Error. (%v + %vi)", resultReal, resultImag)
	}
}

func TestMultiplyBigRatComplexWithSameNumbers(t *testing.T) {
	// (0 + i) * (0 + i) == (-1)
	var (
		aReal = big.NewRat(0, 1)
		aImag = big.NewRat(1, 1)
	)

	resultReal, resultImag := MultiplyBigRatComplex(aReal, aImag, aReal, aImag)
	if resultReal.Cmp(big.NewRat(-1, 1)) != 0 || resultImag.Cmp(big.NewRat(0, 1)) != 0 {
		t.Fatalf("Error. (%v + %vi)", resultReal, resultImag)
	}
}

func TestSquareBigRatComplexWithSameNumbers(t *testing.T) {
	// (0 + i) * (0 + i) == (-1)
	var (
		aReal = big.NewRat(0, 1)
		aImag = big.NewRat(1, 1)
	)

	resultReal, resultImag := SquareBigRatComplex(aReal, aImag)
	if resultReal.Cmp(big.NewRat(-1, 1)) != 0 || resultImag.Cmp(big.NewRat(0, 1)) != 0 {
		t.Fatalf("Error. (%v + %vi)", resultReal, resultImag)
	}
}

func TestSqrtBigRat(t *testing.T) {
	result := SqrtBigRat(big.NewRat(4, 1))
	result.Sub(result, big.NewRat(2, 1))

	if result.Cmp(big.NewRat(1, 10000000000)) != -1 { // should be -1
		t.Fatalf("Error. %v", result)
	}
}
