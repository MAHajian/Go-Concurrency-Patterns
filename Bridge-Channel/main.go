package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Define variables
	input := make(chan int)
	output := make(chan int)

	// Goroutine acts Bridge for communication data between channels
	go func() {
		defer close(output)
		for num := range input {
			output <- num
		}
	}()
	// Goroutine send data to input channel
	go func() {
		defer close(input)
		for i := 0; i < 5; i++ {
			input <- rand.Intn(100)
		}
	}()

	// Get data from output channel
	for num := range output {
		fmt.Println(num)
	}

}
