package main

import "fmt"

func main() {
	// START OMIT
	var ids [3]int
	// ids = nil // compilation error // HL
	fmt.Println(ids) // HL
	// END OMIT
}
