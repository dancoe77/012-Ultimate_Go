package main

import (
	"context"
	"fmt"
	"time"
)

var zz string = "#############################################################"

func main() {
	fmt.Println(zz)
	// Buffered channel to avoid goroutine leak
	ch1, ch2 := make(chan string, 1), make(chan string, 1)

	go func() {
		time.Sleep(100 * time.Millisecond)
		// Possible solution to goroutine leak: select with timeout
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(20 * time.Millisecond) // 200
		ch2 <- "two"
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	select {
	case v := <-ch1:
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2:", v)
		/*
			case <-time.After(10 * time.Millisecond):
				fmt.Println("timeout")
		*/
	case <-ctx.Done():
		fmt.Println("timeout")
	}
	fmt.Println(zz)
}
