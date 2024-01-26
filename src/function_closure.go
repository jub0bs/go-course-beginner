package main

import (
	"fmt"
)

// START OMIT
func adder() func(int) int {
	var sum int // HL
	return func(x int) int {
		sum += x   // HL
		return sum // HL
	}
}

func main() {
	add := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("index: %d; total: %d\n", i, add(i))
	}
}

// END OMIT
