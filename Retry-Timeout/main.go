package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Define variables
	start := time.Now()
	retryTimes := 3
	timeout := time.Duration(time.Second)
	var resp *http.Response
	var err error
	var url string = "https://example.com"

	// Retrying for loop
	for i := 0; i < retryTimes; i++ {
		resp, err = http.Get(url)
		if err == nil {
			break
		}
		fmt.Println("Retrying...")
		<-time.After(timeout)
	}

	// Error handling
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(resp.Status)
	}

	// Calculate elapsed time
	fmt.Println("Elapsed:", time.Since(start))
}
