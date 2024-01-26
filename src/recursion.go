package main

import "fmt"

// START OMIT
func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1) // HL
}

func main() {
	fmt.Println(factorial(5))
}

// END OMIT
