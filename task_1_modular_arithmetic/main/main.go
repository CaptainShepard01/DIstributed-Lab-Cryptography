package main

import (
	"cryptography/task_1_modular_arithmetic/tasks"
	"fmt"
)

func main() {
	var m int
	fmt.Println("Enter module: ")
	_, err := fmt.Scan(&m)
	if err != nil {
		fmt.Println(err)
		return
	}

	a := 25
	firstEquation := tasks.SolveFirstEquation(a, m)
	fmt.Printf("\n%d mod %d is: %d", a, m, firstEquation)

	a = 125
	b := 89
	secondEquation, err := tasks.SolveSecondEquation(a, b, m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\n%d^%d mod %d is: %d", a, b, m, secondEquation)

	a, b = 32, 9
	thirdEquation, err := tasks.SolveThirdEquation(a, b, m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nx in %d*x ≡ %d mod %d is: %d", a, b, m, thirdEquation)

	a, b = 5, 43
	fmt.Printf("\nOne of the prime numbers between %v and %v, is: %v", a, b, tasks.GeneratePrimeNumbersInRange(a, b)[0])
}
