package main

import "fmt"

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	rect := rectangle{width: 5, height: 3}
	fmt.Printf("Area: %.2f\n", rect.area())
	fmt.Printf("Perimeter: %.2f\n", rect.perimeter())

}
