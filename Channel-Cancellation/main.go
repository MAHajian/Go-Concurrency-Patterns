package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	cancel := make(chan bool)
	go doWork(cancel)
	<-time.After(time.Second * 3)
	cancel <- true
	fmt.Println("Elapsed:", time.Since(start))
}

func doWork(cancel <-chan bool) {
	for {
		select {
		case <-cancel:
			fmt.Println("Cancelled")
			return
		default:
			fmt.Println("Working...")
		}
	}
}
