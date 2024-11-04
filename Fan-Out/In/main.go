package main

import "math/rand"

func main() {
	jobs := make(chan int)
	results := make(chan int)
}

func task(jobs<- chan int, results chan<- int) {
	ch <- rand.Intn(100)
}