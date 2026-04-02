package main
import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(sum(5, 3))
	vsum := sum(10, 20)
	fmt.Println("The sum of 10 and 20 is:", vsum)
	// Variables and Constants 
	var x, y int = 15, 25
	result := sum(x, y)
	const pi = 3.14
	fmt.Printf("The value of pi is: %.2f\n", pi)
	fmt.Printf("The sum of %d and %d is: %d\n", x, y, result)

	// pi = 3.14159 // This will cause a compile-time error because pi is a
	// Operators
	a := 10
	b := 3
	fmt.Printf("a + b = %d\n", a + b)
	fmt.Printf("a - b = %d\n", a - b)
	fmt.Printf("a * b = %d\n", a * b)
	fmt.Printf("a / b = %d\n", a / b)
	fmt.Printf("a %% b = %d\n", a % b) // Modulo operator
	c := "Hello, World!"
	fmt.Println(c)
	d := "Go is great!"
	fmt.Println(d)
	fmt.Printf("Concatenated string: %s\n", c + " " + d)
}