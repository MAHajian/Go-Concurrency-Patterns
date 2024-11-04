package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	urls := []string {"url1", "url2", "url3"}
	results := make(chan string)

	// Fan-Out Goroutines
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- downloadFile(url)
		}()
	}

	// Fan-In 
	go func() {
		wg.Wait()
		close(results)
	}()

	var output string
	for result := range results {
		output += result
	}
	fmt.Printf("%v| Elepsed: %v", output, time.Since(start))
}

func downloadFile(url string) string {
	<-time.After(time.Second)
	return fmt.Sprintf("%v:%v ", url,rand.Intn(100))
}