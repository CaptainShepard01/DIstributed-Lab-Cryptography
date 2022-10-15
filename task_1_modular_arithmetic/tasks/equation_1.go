package tasks

import "errors"

// a mod m = x

func SolveFirstEquation(a, m int) (int, error) {
	if m <= 0 {
		return -1, errors.New("invalid module")
	}
	if a < 0 {
		return a - (m * (a / m)) + m, nil
	}
	return a - (m * (a / m)), nil
}
