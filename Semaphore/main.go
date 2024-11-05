package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	SemCh chan struct{}
}

func (s *semaphore) Release() {
	<-s.SemCh
}

func (s *semaphore) Acquire() {
	s.SemCh <- struct{}{}
}

func New(maxConcurrency int) Semaphore {
	return &semaphore{SemCh: make(chan struct{}, maxConcurrency)}
}

func main() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	semaphore := New(1000)
	totalProcess := 10000

	for i := 1; i <= totalProcess; i++ {
		semaphore.Acquire()
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			defer semaphore.Release()
			task(id)
		}(i)
	}
	wg.Wait()
	fmt.Println("Elapsed:", time.Since(start))
}

func task(id int) {
	fmt.Println(time.Now(), "Doing #", id)
	<-time.After(time.Second)
}
