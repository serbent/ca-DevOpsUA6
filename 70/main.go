package main
import "fmt"

func main() {
    fmt.Println("Hello, and welcome to loops conditions and slices!")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
			continue
		} 
		fmt.Printf("%d is odd\n", i)
	}


	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking the loop at 5")
			break
		}
		fmt.Printf("%d ", i)
	}


	name := "GoLang"

	if name == "GoLang" {
		fmt.Println("Welcome to GoLang!")
	} else if name == "Python" {
		fmt.Println("Welcome to Python!")
	} else {
		fmt.Println("Unknown language")
	}


	switch name {
	case "GoLang":
		fmt.Println("Welcome to GoLang!")
	case "Python":
		fmt.Println("Welcome to Python!")
	default:
		fmt.Println("Unknown language")
	}
	fmt.Printf("\n\n\n\n")

	fmt.Println("====  Working with slices in Go  ====")
	
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", slice)

	slice = append(slice, 6)
	fmt.Println("After appending 6:", slice)

	slice = append(slice, 7, 8, 9)
	fmt.Println("After appending 7, 8, 9:", slice)

	slice = append(slice[:2], slice[3:]...)
	fmt.Println("After removing the element at index 2:", slice)

	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	cars := []string{"Toyota", "Honda", "Ford", "BMW", "Audi", "Mercedes", "Volkswagen", "Nissan", "Hyundai", "Kia", "Mazda", "Subaru", "Tesla", "Volvo", "Jaguar", "Land Rover", "Porsche", "Ferrari", "Lamborghini", "Maserati", "Bentley"}
	fmt.Println("Original cars slice:", cars)

	for i := 0; i < len(cars); i++ {
		if cars[i] == "BMW" {
			cars = append(cars[:i], cars[i+1:]...)
			break
		}
}