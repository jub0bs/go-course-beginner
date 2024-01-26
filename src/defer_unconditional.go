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
