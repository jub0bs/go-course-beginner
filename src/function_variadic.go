package main

import (
	"fmt"
)

// START OMIT
func main() {
	fmt.Println(Max(99, 1_000_000, -32))
}

func Max(first int, rest ...int) int { // HL
	m := first
	for _, i := range rest {
		if i > m {
			m = i
		}
	}
	return m
}

// END OMIT
