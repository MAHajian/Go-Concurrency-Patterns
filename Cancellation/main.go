package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	ch := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("Cancelled", "| Elapsed:", time.Since(start))
				return
			case <-time.After(time.Second):
				fmt.Println("Running")
			}
		}
	}()
	<-time.After(time.Second * 3)
	close(ch)
	wg.Wait()
}
