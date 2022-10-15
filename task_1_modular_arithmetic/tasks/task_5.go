package tasks

import (
	"errors"
	"math"
)

func GeneratePrimeNumberInRange(a, b int) (int, error) {
	if a < 2 || b < 2 || a > b {
		return -1, errors.New("wrong bounds")
	}

	for i := a; i <= b; i++ {
		if i == 1 || i == 0 {
			continue
		}

		flag := true

		upperLimit := int(math.Sqrt(float64(i)))
		for j := 2; j <= upperLimit; j++ {
			iModJ, _ := SolveFirstEquation(i, j)
			if iModJ == 0 {
				flag = false
				break
			}
		}

		if flag {
			return i, nil
		}
	}
	return -1, errors.New("no primes in bounds")
}
