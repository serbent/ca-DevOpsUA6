package main

import "fmt"

func main() {
	fmt.Println("Hello, World! Let's start with channels in Go.")

	fmt.Println("!!! Creating unbuffered channel to comunicate between goroutines !!!")
	chan0 := make(chan string)

	go func() {
		fmt.Println("Hello from the goroutine!")
		chan0 <- "Message from goroutine to main function."
	}()

	fmt.Println(<-chan0)

	fmt.Println("!!! Creating buffered channel to communicate between goroutines !!!")
	chan1 := make(chan string, 3)

	chan1 <- "Message 1 from main function."
	chan1 <- "Message 2 from main function."
	chan1 <- "Message 3 from main function."
	fmt.Println(<-chan1)
	fmt.Println(<-chan1)
	fmt.Println(<-chan1)

	// chan1 <- "This is another message from the main function."

	// fmt.Println(<-chan1)
	// fmt.Println(<-chan1)
}
