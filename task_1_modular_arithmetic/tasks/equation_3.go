package tasks

import (
	"errors"
)

// a*x ≡ b mod m

// In arithmetic and computer programming, the extended Euclidean algorithm is an extension to the Euclidean
// algorithm, and computes, in addition to the greatest common divisor (gcd) of integers a and b,
// also the coefficients of Bezout's identity, which are integers x and y such that
// ax + by = gcd(a, b).
//
//	Given two integers ``b`` and ``n``
//
// returns “(gcd(b, n), a, m)“ such that “a*b + n*m == gcd(b, n)“.
func egcd(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	g, y, x := egcd(b%a, a)
	bDiva := (b/a)
	mid := x - bDiva * y
	return g, mid, y
}

// In mathematics, particularly in the area of arithmetic, a modular multiplicative inverse of
// an integer `a` is an integer x such that the product ax is congruent to 1 with respect to the modulus m.
//
//	In the standard notation of modular arithmetic this congruence is written as
//
// ax ≡ 1 (mod m)
// modInv returns multiplicative inverse of an integer a
func findInverseElement(a, m int) (int, error) {
	a = a % m
	g, x, _ := egcd(a, m)
	if g != 1 {
		return -1, errors.New("modular inverse does not exist")
	} else {
		result := x % m
		if result < 0 {
			result += m
		}
		return result, nil
	}
}

func SolveThirdEquation(a, b, m int) (int, error) {
	inverseToA, err := findInverseElement(a, m)
	if err != nil {
		return -1, err
	}

	return (inverseToA * b) % m, nil
}
