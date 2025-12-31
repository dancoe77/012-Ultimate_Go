package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// What are the N most common words in sherlock.txt

// `a` is "raw" string, at \ is just a \
var wordRE = regexp.MustCompile(`[a-zA-Z]+`)

// Code that runs main
// - var expression
// - init(){}

func main() {
	// mapDemo()
	file, err := os.Open("/home/dan/Code/012-Ultimate_Go/04-Panics_and_Maps/15-Calculating_Word_Frequency_and_working_with_Maps/freq/sherlock.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	freq := make(map[string]int) // word -> count
	s := bufio.NewScanner(file)
	// nLines := 0
	for s.Scan() {
		// nLines++
		words := wordRE.FindAllString(s.Text(), -1)
		for _, word := range words {
			freq[strings.ToLower(word)]++
		}

	}
	if err := s.Err(); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	// fmt.Println(nLines)
	top := topN(freq, 10)
	fmt.Println(top)
}

// topN returns the "n" most common words in freq
func topN(freq map[string]int, n int) []string {
	words := slices.Collect(maps.Keys(freq))
	sort.Slice(words, func(i, j int) bool {
		wi, wj := words[i], words[j]
		// Sort in reverse order
		return freq[wi] > freq[wj]
	})

	n = min(n, len(words))
	return words[:n]
}

func mapDemo() {
	heros := map[string]string{ // hero -> name
		"Superman":     "Clark",
		"Wonder Woman": "Diana",
		"Batman":       "Bruce",
	}
	// Keys
	for k := range heros {
		fmt.Println(k)
	}
	fmt.Println("##########################################")

	// Key + Value
	for k, v := range heros {
		fmt.Println(v, "is known as", k)

	}
	fmt.Println("##########################################")

	// for values, use _
	for _, v := range heros {
		fmt.Println(v)

	}
	fmt.Println("##########################################")

	n := heros["Batman"]
	fmt.Println(n)
	fmt.Println("##########################################")

	// Accessing non-existing key it will return the zero value
	n = heros["Aquaman"]
	fmt.Printf("%q\n", n)
	fmt.Println("##########################################")

	// Use "comma" ok idiom to find if a key is in the map
	n, ok := heros["Aquaman"]
	if ok {
		fmt.Printf("%q\n", n)
	} else {
		fmt.Println("Aquaman not found")
	}
	fmt.Println("##########################################")
	delete(heros, "Batman")
	fmt.Println(heros)
	fmt.Println("##########################################")
}
