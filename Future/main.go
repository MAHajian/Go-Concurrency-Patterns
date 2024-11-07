package main

import (
	"fmt"
	"math/rand"
	"time"
)

type future struct {
	result chan int
}

func (f *future) Get() int {
	return <-f.result
}

func task() future {
	f := &future{make(chan int)}
	go func() {
		<-time.After(time.Second)
		f.result <- rand.Intn(100)
	}()
	return *f
}

func main() {
	f := task()
	fmt.Println(f.Get())
}
