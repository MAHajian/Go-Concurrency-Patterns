package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	done := make(chan bool)
	go task(done)
	<-done
	fmt.Println("Task Complete!", "| Elapsed:",time.Since(start))
}

func task(done chan<- bool) {
	fmt.Println("Doing Task...")
	<-time.After(time.Second)
	done <- true
}