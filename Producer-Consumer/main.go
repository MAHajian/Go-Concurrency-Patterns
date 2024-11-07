package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go producer(ch, time.Millisecond*200)
	for i := 1; i <= 2; i++ {
		go consumer(i, ch, time.Millisecond*100)
	}
	<-time.After(time.Second)
}

func producer(ch chan<- int, duration time.Duration) {
	for {
		data := rand.Intn(100)
		ch <- data
		fmt.Println("Producer:", data)
		<-time.After(duration)
	}
}

func consumer(id int, ch <-chan int, duration time.Duration) {
	for data := range ch {
		fmt.Printf("Customer # %v Received: %v\n", id, data)
		<-time.After(duration)
	}
}
