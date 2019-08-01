package main

import "math/big"

// big.Float

// MultiplyBigFloatComplex returns two values, real part and imaginary part
// xy = (a + ib)(c + ib) = (ac - bd) + i(ad + bc)
func MultiplyBigFloatComplex(aReal, aImag, bReal, bImag *big.Float) (*big.Float, *big.Float) {
	var r, i, a, b big.Float
	return r.Sub(a.Mul(aReal, bReal), b.Mul(aImag, bImag)), i.Add(a.Mul(aReal, bImag), b.Mul(aImag, bReal))
}

// SquareBigFloatComplex calculate square of complex number that is represented by big.Float
// and return two values, real part and imaginary part
func SquareBigFloatComplex(vReal, vImag *big.Float) (*big.Float, *big.Float) {
	return MultiplyBigFloatComplex(vReal, vImag, vReal, vImag)
}

// AbsBigFloatComplex calculate absolute of complex number that is represented by big.Float
// and return two values, real part and imaginary part
func AbsBigFloatComplex(vReal, vImag *big.Float) *big.Float {
	var r, a, b big.Float
	r.Add(a.Mul(vReal, vReal), b.Mul(vImag, vImag))
	return r.Sqrt(&r)
}

// big.Rat

// MultiplyBigRatComplex returns two values, real part and imaginary part
// xy = (a + ib)(c + ib) = (ac - bd) + i(ad + bc)
func MultiplyBigRatComplex(aReal, aImag, bReal, bImag *big.Rat) (*big.Rat, *big.Rat) {
	var r, i, a, b big.Rat
	return r.Sub(a.Mul(aReal, bReal), b.Mul(aImag, bImag)), i.Add(a.Mul(aReal, bImag), b.Mul(aImag, bReal))
}

// SquareBigRatComplex calculate square of complex number that is represented by big.Rat
// and return two values, real part and imaginary part
func SquareBigRatComplex(vReal, vImag *big.Rat) (*big.Rat, *big.Rat) {
	return MultiplyBigRatComplex(vReal, vImag, vReal, vImag)
}

// AbsBigRatComplex calculate absolute of complex number that is represented by big.Rat
// and return two values, real part and imaginary part
func AbsBigRatComplex(vReal, vImag *big.Rat) *big.Rat {
	var r, a, b big.Rat
	r.Add(a.Mul(vReal, vReal), b.Mul(vImag, vImag))
	return SqrtBigRat(&r)
}

// SqrtBigRat find suare root of big.Rat and return big.Rat.
// > Finding sqrt(S) is the same as solving the equation f(x) = x^2 - S = 0 for a positive x.
//  https://en.wikipedia.org/wiki/Methods_of_computing_square_roots
//  Finding sqrt(S) using newton's method.
//  x_(n+1) = x_n - (f(x_n) / f'(x_n))
//          => x - (x^2 - S) / 2x
// 			= (x + S/x) / 2
func SqrtBigRat(s *big.Rat) *big.Rat {
	const iterations = 10

	x := big.NewRat(10, 1)  // temp
	tmp := new(big.Rat)     // temp
	two := big.NewRat(2, 1) // const
	for n := uint8(0); n < iterations; n++ {
		tmp.Quo(s, x)   // S/x
		tmp.Add(x, tmp) // x + S/x
		x.Quo(tmp, two) // (x + S/x) / 2
	}
	return x
}
