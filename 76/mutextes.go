package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	counter := 0

	for i := 0; i < 5; i++ {
		go func() {
			mutex.Lock()
			counter++
			defer mutex.Unlock()

			fmt.Printf("Counter: %d\n", counter)
		}()
	}

	// Wait for user input to allow goroutines to finish
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
