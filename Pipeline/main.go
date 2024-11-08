package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	nums := []int{5, 2, 10, 15}
	ch := gen(nums...)
	output := sq(ch)
	for out := range output {
		fmt.Println(out)
	}
	fmt.Println("Elapsed:", time.Since(start))
}

func gen(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, num := range nums {
			ch <- num
		}
	}()
	return ch
}

func sq(nums <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for num := range nums {
			ch <- (num * num)
		}
	}()
	return ch
}
