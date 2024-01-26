package main

import (
	"fmt"
	"strconv"
)

func main() {
	// START OMIT
	str := "42.2"
	if _, err := strconv.Atoi(str); err != nil { // HL
		fmt.Printf("%q cannot be parsed to an integer: %v\n", str, err)
	}
	str = "89"
	if _, err := strconv.Atoi(str); err != nil { // HL
		fmt.Printf("%q cannot be parsed to an integer: %v\n", str, err)
	}
	// END OMIT
}
