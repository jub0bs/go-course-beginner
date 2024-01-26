package main

import (
	"fmt"
	"unicode/utf8"
)

// START OMIT
func main() {
	hello := "Hello, "
	world := "世界"
	s := hello + world
	fmt.Printf("%q\n", s)
	fmt.Printf("number of bytes: %d\n", len(s))
	fmt.Printf("number of runes: %d\n", utf8.RuneCountInString(s))
}

// END OMIT
