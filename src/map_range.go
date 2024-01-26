package main

import (
	"fmt"
)

func main() {
	// START OMIT
	m := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}
	for k, v := range m { // HL
		fmt.Printf("%s: %q\n", k, v)
	}
	// END OMIT
}
