package main

import (
	"fmt"
	"time"
)

var l string = "##############################################################"

func main() {
	fmt.Println(l)
	go fmt.Println("goroutine")
	fmt.Println(l)
	fmt.Println("main")
	fmt.Println(l)

	for i := range 3 {
		// Prior to Go 1.22 this was a bug
		go func() {
			fmt.Println("goroutine:", i)
		}()
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println(l)
	ch := make(chan int)
	go func() {
		ch <- 7 // send
	}()
	v := <-ch // receive
	fmt.Println(v)
	fmt.Println(l)

	fmt.Println(sleepSort([]int{20, 30, 10})) //[10 20 30]
	fmt.Println(l)

	go func() {
		for i := range 4 {
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(">>", v)
	}
	fmt.Println(l)
	v = <-ch // ch is closed
	fmt.Println("closed:", v)
	fmt.Println(l)

	v, ok := <-ch // ch is closed
	fmt.Println("closed:", v, "ok:", ok)
	fmt.Println(l)

	/*
		The "for range" above does
		for {
			v, ok := <- ch
			if !ok {
				break
			}
			fmt.Println(">>", v)
		}
	*/

	// var ch chan int // ch is nil
}

/*
Channel semantics
- send/receive to/from a channel will block until opposite operation(*)
	- guarantee of delivery
- receive from a closed channel will return the zero value without blocking
	- use "comma ok" idiom to check if channel was closed
- send to a closed channel will cause a panic
- closing a closed or nil channel will cause a panic
- send/receive to a nil channel will block forever
*/

/*
Algorithm
- For every value "n" in values, spin a goroutine that
  - sleeps for "n" milliseconds
  - sends "n" over a channel

- collect all values from the channel to a slice and return it
*/
func sleepSort(values []int) []int {
	ch := make(chan int)

	for _, n := range values {
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}()
	}
	var out []int
	for range values {
		n := <-ch
		out = append(out, n)
	}
	return out
}
