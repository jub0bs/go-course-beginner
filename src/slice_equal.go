package main

import (
	"fmt"
	"slices"
)

func main() {
	// START OMIT
	foo := []int{42, 56, 70}
	bar := []int{98, 99}
	fmt.Println(slices.Equal(foo, bar)) // HL
	// END OMIT
}
