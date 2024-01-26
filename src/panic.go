package main

import (
	"fmt"
)

// START OMIT
func main() {
	fmt.Println("Hello, World!")
	panic("oops") // HL
	fmt.Println("Goodbye, World!")
}

// END OMIT
