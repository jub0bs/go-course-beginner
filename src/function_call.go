package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	const (
		gandalfWisdom = "All you have to do is to decide what to do with the time given to you."
		sep           = " "
	)
	words := strings.Split(gandalfWisdom, sep) // HL
	fmt.Printf("%q\n", words)
}
// END OMIT

