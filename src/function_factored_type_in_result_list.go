package main

import (
	"fmt"
)

// START OMIT
func main() {
	fmt.Println(MinAndMax(99, 1_000_000, -32))
}

func MinAndMax(first int, rest ...int) (min, max int) { // HL
	// omitted
	// END OMIT
	min, max = first, first
	for _, i := range rest {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return
}
