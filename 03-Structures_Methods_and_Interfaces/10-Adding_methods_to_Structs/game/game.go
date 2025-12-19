package main

import (
	"fmt"
	"slices"
)

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

	i.Move(10, 20)
	fmt.Printf("i (move): %#v\n", i)

	p1 := Player{
		Name: "Parzival",
	}
	fmt.Printf("p1: %+v\n", p1)
	fmt.Println("p1.X:", p1.X)
	p1.Move(100, 200)
	fmt.Printf("p1 (move): %+v\n", p1)

	fmt.Println(p1.Found("copper")) // <nil>
	fmt.Println(p1.Found("copper")) // <nil>
	fmt.Println(p1.Found("gold"))   // unknown key: "gold"
	fmt.Println("keys:", p1.Keys)   // keys: [copper]
}

/*
Exercise:
- Add a "Keys" field to Player which is a slice of strings
- Add a "Found(key string)" method to player
	- It should err if key is not one of "jade", "copper", "crystal"
	- It should add a key only once
*/

func (p *Player) Found(key string) error {
	switch key {
	case "copper", "crystal", "jade":
		// Ok
	default:
		return fmt.Errorf("unknown key: %q\n", key)
	}
	if !slices.Contains(p.Keys, key) {
		p.Keys = append(p.Keys, key)
	}
	return nil
}

type Player struct {
	Name string
	Item // Player embeds Item
	Keys []string
}

// Move will change the value of i by delta x and delta y
// "i" is called "the receiver"
// "i" is a pointer receiver
/*
Value vs pointer receiver
- In general it is recommended to use Value semantics
- Try to keep same semantics on all methods
- When you must use pointer semantics
	- If you have a lock field, ie Mutex or Atomic
	- If you need to mutate the struct
	- Decoding/unmarshalling
*/
func (i *Item) Move(dx, dy int) {
	i.X += dx
	i.Y += dy
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
