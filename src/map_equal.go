package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}
	// START OMIT
	m2 := map[string]string{"white": "ffffff"}
	fmt.Println(maps.Equal(m, m2)) // HL
	// END OMIT
}
