package main

import "fmt"

func main() {
	var a any // interface{}

	a = 7
	fmt.Println("a:", a)

	a = "Hi"
	fmt.Println("a:", a)
	/*
		Rule of thumb: Don't use any :)
		Exceptions:
		- Serialization
		- Printing
	*/

	s := a.(string) // type assertion
	fmt.Println("s:", s)

	// i := a.(int) // Will panic
	// AKA: "comma, ok" idiom; "ok" is a bool
	i, ok := a.(int)
	if ok {
		fmt.Println("i:", i)
	} else {
		fmt.Printf("not an int (%T)\n", a)
	}

	// type switch
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Printf("other type (%T)\n", a)
	}
}
