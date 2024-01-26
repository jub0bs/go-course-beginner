package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func heads() bool {
	return rand.Intn(2) == 0 // true half the time
}

func main() {
	// START OMIT
	for {
		if heads() { // HL
			fmt.Println("Heads")
		} else {
			fmt.Println("Tails")
			break // HL
		}
	}
	// END OMIT
}
