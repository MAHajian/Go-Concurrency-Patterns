package main

import (
	"fmt"
	"math/rand"
)

func main() {
	in := make(chan int, 10)
	out := make(chan int, 10)
	data := []int{}
	
	go func() {
		for {
			select {
			case in <- rand.Intn(100):
			default:
				fmt.Println("Dropping input data")
			}
		}
	}()

	go func() {
		for {
			select {
			case data := <- in:
				select {
				case out <- data:
				default:
					fmt.Println("Dropping output data")
				}
			}
		}
	}()

	for i := 0; i < 11; i++ {
		data = append(data, <-out)
	}
	fmt.Println(data)
}