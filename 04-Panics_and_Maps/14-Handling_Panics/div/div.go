package main

import "fmt"

func main() {
	fmt.Println(safeDiv(7, 3))
	fmt.Println(safeDiv(2, 0))
}

/*
Using named return values:
- defer/recover to change return error value
- documentation
*/

func safeDiv(a, b int) (q int, err error) {
	// q & err are variables inside safeDiv
	// just like a & b
	defer func() {
		if e := recover(); e != nil {
			//fmt.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	return div(a, b), nil
}

func div(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
