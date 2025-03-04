package main

import (
	"fmt"
	"math/rand/v2"
)

func heads() bool {
	return rand.Intn(2) == 0 // true half the time
}

func main() {
	// START OMIT
	for {
		if heads() { // HL
			fmt.Println("Heads")
			continue // HL
		}
		fmt.Println("Tails")
		return
	}
	// END OMIT
}
