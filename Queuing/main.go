package main

import (
	"fmt"
	"math/rand"
)

func main() {
	queue := make(chan int, 5)

	go func() {
		defer close(queue)
		for i := 0; i < 5; i++ {
			data := rand.Intn(100)
			queue <- data
			fmt.Println("Sent:", data)
		}
	}()

	for q := range queue {
		fmt.Println("Received:", q)
	}
}
