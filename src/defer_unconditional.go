package main

import (
	"fmt"
	"math/rand/v2"
)

func heads() bool {
	return rand.IntN(2) == 0 // true half the time
}

// START OMIT
func main() {
	defer fmt.Println("Gambling is bad for your health.")
	if heads() {
		fmt.Println("You win!")
	} else {
		fmt.Println("You lose...")
		panic("😱")
	}
}

// END OMIT
