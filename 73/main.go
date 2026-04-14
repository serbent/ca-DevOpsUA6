package main

import "fmt"

func main() {

	fmt.Println("Hello, from pointers and recerivers presentaption")

	TheMeaningOfLive := 42
	fmt.Println("The meaning of live is", TheMeaningOfLive)

	AddressOfTheMeaningOfLive := &TheMeaningOfLive
	fmt.Println("The address of the meaning of live is", AddressOfTheMeaningOfLive)

	var PointerToTheMeaningOfLive *int = &TheMeaningOfLive
	fmt.Println("The pointer to the meaning of live is", PointerToTheMeaningOfLive)

	var ValueOfTheMeaningOfLive int = *PointerToTheMeaningOfLive
	fmt.Println("The value of the meaning of live is", ValueOfTheMeaningOfLive)

	var p *int
	p = &TheMeaningOfLive
	fmt.Println("The value of p is", *p)

	*p = 21
	fmt.Println("The value of TheMeaningOfLive is", TheMeaningOfLive)
}
