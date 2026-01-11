package main

import (
	"fmt"
	"sync"
	"time"
)

var zz string = "############################################################################"

func main() {
	fmt.Println(zz)
	/*
		Solution 1: mutex
	*/
	var mu sync.Mutex
	count := 0
	/*
		Solution 2: sync/atomic //avoid unless absolutely necessary
		count := int64(0)
	*/
	nGR, nIter := 10, 1_000

	var wg sync.WaitGroup

	wg.Add(nGR)
	for range nGR {
		go func() {
			defer wg.Done()
			for range nIter {
				/*
					atomic.AddInt64(&count, 1)
				*/
				mu.Lock()
				count++
				mu.Unlock()
				/*
					fetch count
					increment count
					store count
				*/
				time.Sleep(time.Microsecond)
			}
			fmt.Println(zz)
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)
	fmt.Println(zz)
}

/*
go run -race count.go

"-race" is supported by
- run
- build
- test
*/
