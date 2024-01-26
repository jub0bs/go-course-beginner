package main

import (
	"fmt"
)

func main() {
	// START OMIT
	s1 := "\u00e9"
	s2 := "e\u0301"
	fmt.Printf("%[1]q: %[1]x\n", s1)
	fmt.Printf("%[1]q: %[1]x\n", s2)
	// END OMIT
}
