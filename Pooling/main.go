package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(w, jobs, results)
		}()
	}

	for job := 0; job < 5; job++ {
		jobs <- rand.Intn(100)
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker #", id, "doing", job)
		results <- job * 2
	}
}
