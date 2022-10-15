package tasks

import (
	"errors"
)

// a*x ≡ b mod m

func egcd(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	bModA, _ := SolveFirstEquation(b, a)
	g, y, x := egcd(bModA, a)
	bDiva := b / a
	mid := x - bDiva*y
	return g, mid, y
}

func findInverseElement(a, m int) (int, error) {
	a = a % m
	g, x, _ := egcd(a, m)
	if g != 1 {
		return -1, errors.New("modular inverse does not exist")
	} else {
		result, _ := SolveFirstEquation(x, m)
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

	result, _ := SolveFirstEquation(inverseToA*b, m)
	return result, nil
}
