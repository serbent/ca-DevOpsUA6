package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func sumAll(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	a := 10
	b := 20
	fmt.Printf("==== A function to sum integers ====\n")
	result := sum(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, result)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, sum(a, b))
	fmt.Println("Hello, World!", "another message")

	fmt.Printf("==== A function to minus integers ====\n")
	result = minus(a, b)
	fmt.Printf("The minus of %d and %d is %d\n", a, b, result)
	fmt.Printf("The minus of %d and %d is %d\n", a, b, minus(a, b))

	fmt.Printf("==== A function to sum all integers with variadic parameters ====\n")
	result = sumAll(a, b, 30, 40, 50)
	fmt.Printf("The sum of %d, %d, %d and %d is %d\n", a, b, 30, 40, result)
	fmt.Printf("The sum of %d, %d, %d and %d is %d\n", a, b, 30, 40, sumAll(a, b, 30, 40, 50, 60))

	fmt.Printf("==== A function to calculate factorial ====\n")
	n := uint64(6)
	fmt.Printf("The factorial of %d is %d\n", n, factorial(n))
	// fmt.Printf("The factorial of %d is %d\n", n, factorial(n))
}
