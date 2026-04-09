package main

import (
	"fmt"
	"math/big" // big package is used to handle large integers, such as factorial of 99
)

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

func factorialBig(n uint64) *big.Int {
	result := big.NewInt(1)
	for i := uint64(2); i <= n; i++ {
		result.Mul(result, new(big.Int).SetUint64(i))
	}
	return result
}

func factorial(n int) int {
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

	fmt.Printf("==== A function to calculate big factorial ====\n")
	n := uint64(99)
	fmt.Printf("The factorial of %d is %d\n", n, factorialBig(n))
	// fmt.Printf("The factorial of %d is %d\n", n, factorial(n))

	fmt.Printf("==== A function to calculate simple factorial ====\n")
	m := 9
	fmt.Printf("The factorial of %d is %d\n", m, factorial(m))
}
