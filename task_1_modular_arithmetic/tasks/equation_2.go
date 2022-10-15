package tasks

import (
	"errors"
	"fmt"
)

// a^b mod m = x

func getBinaryRepresentation(number int) string {
	return fmt.Sprintf("%b", number)
}

func SolveSecondEquation(a, b, m int) (int, error) {
	if a == 0 {
		if b == 0 {
			return -1, errors.New("0^0 is undefined")
		}
		return 0, nil
	}
	if m == 0 {
		return -1, errors.New("module is 0")
	}
	d := 1
	t := a
	runes := []rune(getBinaryRepresentation(b))

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == '1' {
			d, _ = SolveFirstEquation(d*t, m)
		}

		t, _ = SolveFirstEquation(t*t, m)
		if t == 1 {
			break
		}
	}

	return d, nil
}
