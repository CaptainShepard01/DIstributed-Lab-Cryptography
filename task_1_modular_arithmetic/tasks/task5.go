package tasks

func GeneratePrimeNumbersInRange(a, b int) []int {
	result := make([]int, 0)
	for i := a; i <= b; i++ {
		if i == 1 || i == 0 {
			continue
		}

		flag := true

		for j := 2; j <= i / 2; j++ {
			if i % j == 0 {
				flag = false
				break
			}
		}

		if flag {
			result = append(result, i)
		}
	}
	return result
}