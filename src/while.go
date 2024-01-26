package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START OMIT
func heads() bool {
	return rand.Intn(2) == 0 // true half the time
}

func main() {
	for heads() { // HL
		fmt.Println("Heads")
	}
	fmt.Println("Tails")
}

// END OMIT
