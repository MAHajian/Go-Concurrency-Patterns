package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Declare Variables
	start := time.Now()
	wg := new(sync.WaitGroup)
	tasks := make(chan int, 100)
	var numOfWorkers, numOfTasks int = 5, 20

	// Initialize Workers & Doing Tasks By Workers
	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for task := range tasks {
				doTask(id, task)
			}
		}(i)
	}

	// Send Tasks to Channel
	for i := 0; i < numOfTasks; i++ {
		tasks <- rand.Intn(100)
	}

	// Close Channel
	close(tasks)
	wg.Wait()
	
	// Calculate Elaspsed Time From After User
	fmt.Println("Elapsed:", time.Since(start))
}

// Doing Task Function
func doTask(id int, task int) {
	fmt.Println("Goroutine #", id, "doing:", task)
	<-time.After(time.Second)
}
