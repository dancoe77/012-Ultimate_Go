package main

import "fmt"

type Item struct {
	X int
	Y int
}

const (
	maxX = 600
	maxY = 400
)

func main() {
	var i Item
	fmt.Printf("i: %#v\n", i)

	i = Item{10, 20} // must specify all fields
	fmt.Printf("i: %#v\n", i)

	i = Item{
		X: 11,
		Y: 22,
	}
	fmt.Printf("i: %#v\n", i)

	// Can be zero value
	i = Item{
		//X: 11,
		Y: 22,
	}
	fmt.Printf("i: %#v\n", i)

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, 2000))

	fmt.Println(NewItemPointer(10, 20))
	fmt.Println(NewItemPointer(10, 2000))

	/* Aside: Use %#v for debugging/logging
	a, b := 1, "1"
	fmt.Printf("a=%v, b=%v\n", a, b)
	fmt.Printf("a=%#v, b=%#v\n", a, b)
	*/
}

/* Types of "new" or factory functions
func NewItem(x, y int) Item
func NewItem(x, y int) *Item
func NewItem(x, y int) (Item, error)
func NewItem(x, y int) (*Item, error)

Value semantics: everyone has their own copy
Pointer semantics: everyone shares the same copy (heap, lock)
*/

func NewItem(x, y int) (Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		// Value Semantics
		return Item{}, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}
	i := Item{
		X: x,
		Y: y,
	}
	return i, nil
}

func NewItemPointer(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		// Value Semantics
		// return Item{}, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}
	i := Item{
		X: x,
		Y: y,
	}
	// The go compiler does escape analysis and will allocate i on the heap
	return &i, nil
}
