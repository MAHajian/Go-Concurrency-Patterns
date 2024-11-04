package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan int)
	go task(ch)
	fmt.Println("result:", <-ch,"- elapsed time:", time.Since(start))
}

func task(ch chan<- int) {
	// doing task
	<-time.After(time.Second)
	ch <- 1
}