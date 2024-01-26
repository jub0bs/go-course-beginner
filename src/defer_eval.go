package main

import (
	"fmt"
)

// START OMIT
func main() {
	var n int

	defer func(i int) { // HL
		fmt.Println("defer with param:", i)
	}(n)

	n++
	defer func() { // HL
		fmt.Println("defer without param:", n)
	}()

	n++
	fmt.Println("normal print", n)
}

// END OMIT
