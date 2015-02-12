package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("0x%x", "世界"[1]) // => 0xb8
	fmt.Println("")
	fmt.Println(len("Hello, 世界")) // => 13
	fmt.Println(len("ĄŚżę"))      // => 8

	s := "Hello, 世界"
	for _, r := range s {
		fmt.Printf("%c ", r) // => H e l l o ,   世 界
	}
	fmt.Println("")

	fmt.Println(utf8.RuneCountInString(s)) // => 9
	fmt.Println(utf8.ValidString(s))       // => true
}
