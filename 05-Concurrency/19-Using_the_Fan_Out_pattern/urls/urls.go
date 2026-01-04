package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

var l string = "###############################################"

func main() {
	fmt.Println(l)
	urls := []string{
		"https://go.dev",
		"https://ardanlabs.com",
		"https://ibm.com/no/such/page",
	}
	start := time.Now()
	/*
		for _, url := range urls {
			stat, err := urlCheck(url)
			fmt.Printf("%q: %d (%v)\n", url, stat, err)
		}
	*/
	fanOutResult(urls)
	duration := time.Since(start)
	fmt.Printf("%d urls in %v\n", len(urls), duration)
	fmt.Println(l)

	start_wait := time.Now()
	fanOutWait(urls)
	duration_wait := time.Since(start_wait)
	fmt.Printf("%d urls in %v\n", len(urls), duration_wait)
	fmt.Println(l)

	start_pool := time.Now()
	fanOutPool(urls)
	duration_pool := time.Since(start_pool)
	fmt.Printf("%d urls in %v\n", len(urls), duration_pool)
	fmt.Println(l)

}

func fanOutResult(urls []string) {
	type result struct {
		url    string
		status int
		err    error
	}
	ch := make(chan result)
	// fanOut
	for _, url := range urls {
		go func() {
			r := result{url: url}
			defer func() { ch <- r }()

			r.status, r.err = urlCheck(url)
		}()
	}
	// collect results
	for range urls {
		r := <-ch
		fmt.Printf("%q: %d (%v)\n", r.url, r.status, r.err)
	}
}

func fanOutWait(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	// fan-out
	for _, url := range urls {
		// wg.Add(1)
		go func() {
			defer wg.Done()
			urlLog(url)
		}()
	}
	//wait for goroutines to finish
	// If you need errors, check out errgroup
	wg.Wait()
}
func fanOutPool(urls []string) {
	var wg sync.WaitGroup

	ch := make(chan string)

	// Producer
	go func() {
		for _, url := range urls {
			ch <- url
		}
		close(ch)
	}()

	const size = 2
	wg.Add(size)
	for range size {
		// Consumers
		go func() {
			defer wg.Done()
			for url := range ch {
				urlLog(url)
			}
		}()
	}
	//wait for goroutines to finish
	wg.Wait()
}

func urlCheck(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

func urlLog(url string) {
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("urlLog", "url", url, "error", err)
		return
	}
	slog.Info("urlLog", "url", url, "status", resp.StatusCode)
}
