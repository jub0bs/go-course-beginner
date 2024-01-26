package main

import (
	"fmt"
)

// START OMIT
func main() {
	next := map[int]int{1: 2, 2: 3}
	n := len(next)
	for k := range next {
		next[k+n]++ // don't do this! // HL
	}
	fmt.Println(next)
}

// END OMIT
