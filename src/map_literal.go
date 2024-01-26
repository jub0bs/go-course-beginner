package main

import (
	"fmt"
)

func main() {
	// START OMIT
	m := map[string]string{ // HL
		"red":   "#ff0000", // HL
		"blue":  "#0000ff", // HL
		"green": "#00ff00", // this final comma is mandatory // HL
	} // HL
	fmt.Println(m == nil)
	fmt.Println(m)
	// END OMIT
}
