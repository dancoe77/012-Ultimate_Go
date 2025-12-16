package main

import (
	"fmt"
	"sort"
)

func main() {
	cart := []string{"apple", "orange", "banana"}
	fmt.Println("len:", len(cart))
	fmt.Println("cart[1]:", cart[1])

	// indices
	for i := range cart {
		fmt.Println(i)
	}

	// index + value
	for i, c := range cart {
		fmt.Println(i, c)
	}

	// values
	for _, c := range cart {
		fmt.Println(c)
	}

	cart = append(cart, "milk")
	fmt.Println(cart)

	// slicing operator, half-open
	fruit := cart[:3]
	fmt.Println("fruit:", fruit)
	fruit = append(fruit, "lemon")
	fmt.Println("fruit:", fruit)
	fmt.Println("cart:", cart)

	var s []int
	for i := range 10_000 {
		s = appendInt(s, i)
	}
	fmt.Println(s[:10])

	// Exercise: concat, without using a "for" loop
	out := concat([]string{"A", "B"}, []string{"C"})
	fmt.Println("concat", out) // [A B C]

	values := []float64{3, 1, 2}
	fmt.Println(median(values)) // 2
	values = []float64{3, 1, 2, 4}
	fmt.Println(median(values)) // 2.5
	fmt.Println("values:", values)

	players := []Player{
		{"Rick", 10_000},
		{"Morty", 11},
	}

	// Add a bonus
	// Value semantics "for" loop
	for _, p := range players {
		p.Score += 100
	}
	fmt.Println("players:", players)

	// "Pointer" semantics for loop
	for i := range players {
		players[i].Score += 100
	}
	fmt.Println(players)
}

type Player struct {
	Name  string
	Score int
}

/*
median
- sort values
- if odd number of values: return middle
- return average of middles
*/
func median(values []float64) float64 {
	// Copy in order not to mutate input argument
	vals := make([]float64, len(values))
	copy(vals, values)

	sort.Float64s(vals)
	i := len(vals) / 2
	if len(vals)%2 == 1 {
		return vals[i]
	}
	mid := (vals[i-1] + vals[i]) / 2
	return mid
}

func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) == cap(s) {
		// no more space in underlying array
		// need to reallocate and copy
		size := 2 * (len(s) + 1)
		fmt.Println(cap(s), "->", size)
		ns := make([]int, size)
		copy(ns, s)
		s = ns[:len(s)]
	}

	s = s[:len(s)+1]
	s[i] = v
	return s
}
