package main

import (
	"fmt"
)

func main() {
	// START OMIT
	const r = 'Z'
	if 'A' <= r && r <= 'Z' || 'a' <= r && r <= 'z' { // HL
		fmt.Printf("%q is alphabetical ASCII\n", r)
	}
	// END OMIT
}
