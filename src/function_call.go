package main

import (
	"fmt"
	"strings"
)

func main() {
	// START OMIT
	fmt.Printf("%T\n", strings.Repeat) // func(string, int) string
	const greetings = "Greetings, Professor Falken. Shall we play a game?\n"
	// fmt.Println(strings.Repeat())          // compilation error: missing both arguments // HL
	// fmt.Println(strings.Repeat(greetings)) // compilation error: missing int argument // HL
	// fmt.Println(strings.Repeat(3))         // compilation error: missing string argument // HL
	fmt.Println(strings.Repeat(greetings, 3)) // all good // HL
	// END OMIT
}
