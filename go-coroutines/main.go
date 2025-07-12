package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var successful []string

// mutex is like control the memory access to memory for threads or goroutines
var mut sync.Mutex

func main() {
	websites := []string{
		"https://google.com",
		"https://facebook.in",
		"https://github.com",
		"https://instagram.com",
		"https://go.dev",
	}

	for _, url := range websites {
		go getStatusCode(url)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(successful)
}

func getStatusCode(url string) {
	defer wg.Done()

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Failed to get status code")
	} else {
		mut.Lock()
		defer mut.Unlock()
		successful = append(successful, url)
		fmt.Printf("%d status code returned from %s\n", res.StatusCode, url)
	}

}
