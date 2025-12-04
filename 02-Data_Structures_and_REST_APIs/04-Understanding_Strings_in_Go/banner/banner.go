package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G", 6)

	// s := "G"
	// fmt.Println("len:", len(s))
	// fmt.Println("s[1]:", s[1])
	// fmt.Printf("s[1]: &c\n", s[1])

	// for i, c := range s {
	// 		fmt.Printf("%c at %d\n", c, i)
	//}

}

func banner(text string, width int) {
	// BUG: len is in byte
	// padding := (width - len(text)) / 2
	padding := (width - utf8.RuneCountInString(text)) / 2
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(text)
	fmt.Println(strings.Repeat("-", width))
}

// banner("Go", 6)
// 	Go
// ------
/*
strings are UTF-8 encoded
len, s[]: byte (uint8)
for : rune (int32)
*/
