// Reference:
// - https://en.wikipedia.org/wiki/Fibonacci_number
// - http://mathworld.wolfram.com/FibonacciNumber.html
// - http://mathworld.wolfram.com/BinetsFibonacciNumberFormula.html
package main

import (
	"math"
	"math/big"
)

var (
	sqrt5  = math.Sqrt(5)
	phi    = (1 + sqrt5) / 2
	phiNeg = (1 - sqrt5) / 2
	logPhi = math.Log(phi)
)

func round(a float64) int {
	if math.Abs(a) < 0.5 {
		return 0
	}
	return int(a + math.Copysign(0.5, a))
}

// TODO accept big.Int
func cloestFibN(a int) int {
	if a < 0 {
		return 0
	}
	if a < 2 {
		return a
	}

	return int(round(math.Log(float64(a)*sqrt5) / logPhi))
}

type M struct {
	M00, M01, M10, M11 *big.Int
}

func badd(a *big.Int, b *big.Int) *big.Int {
	c := big.NewInt(0)
	c.Add(a, b)
	return c
}

func bmul(a *big.Int, b *big.Int) *big.Int {
	c := big.NewInt(0)
	c.Mul(a, b)
	return c
}

func (a M) Mul(b M) (c M) {
	// TODO improve memory usage
	c.M00 = badd(bmul(a.M00, b.M00), bmul(a.M01, b.M10))
	c.M01 = badd(bmul(a.M00, b.M01), bmul(a.M01, b.M11))
	c.M10 = badd(bmul(a.M10, b.M00), bmul(a.M11, b.M01))
	c.M11 = badd(bmul(a.M10, b.M01), bmul(a.M11, b.M11))

	return
}

func (m M) Pow(n int64) M {
	if n == 1 {
		return m
	}

	halfN := m.Pow(n / 2)
	powN := halfN.Mul(halfN)
	if n%2 == 1 {
		powN = powN.Mul(m)
	}

	return powN
}

var fibSeq = []int64{
	0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987,
	1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025,
	121393, 196418, 317811, 514229, 832040, 1346269,
	2178309, 3524578, 5702887, 9227465, 14930352, 24157817,
	39088169, 63245986, 102334155, 165580141, 267914296, 433494437,
}

func fibAtN(n int) *big.Int {
	if n < 0 {
		return big.NewInt(fibSeq[0])
	}
	if n < len(fibSeq) {
		return big.NewInt(fibSeq[n])
	}

	fibM := M{big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(0)}
	return fibM.Pow(int64(n) - 1).M00
}

func fibNext(n int) string {
	return fibAtN((cloestFibN(n) + 1)).String()
}
