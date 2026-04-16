package main

import "fmt"

type geometry interface {
	area() float64
	perimeter() float64
	printType() string
}

type rectanglev struct {
	width, height float64
}

func (r rectanglev) area() float64 {
	return r.width * r.height
}

func (r rectanglev) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r rectanglev) printType() string {
	return "Rectangle"
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14159 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14159 * c.radius
}

func (c circle) printType() string {
	return "Circle"
}

type triangle struct {
	base, height float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) perimeter() float64 {
	// Assuming an equilateral triangle for simplicity
	return 3 * t.base
}

func (t triangle) printType() string {
	return "Triangle"
}

type square struct {
	side float64
}

func mesure(g geometry) {
	fmt.Println("The struc: ", g)
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perimeter: %.2f\n", g.perimeter())
	fmt.Printf("Type: %s\n", g.printType())
}

func main() {
	rect := rectanglev{width: 5, height: 3}
	circ := circle{radius: 4}
	tri := triangle{base: 5, height: 3}
	squ := square{side: 4}

	mesure(rect)
	fmt.Println("")
	mesure(circ)
	fmt.Println("")
	mesure(tri)
	fmt.Println("")
	mesure(squ) // Error as square does not implement geometry interface
}
