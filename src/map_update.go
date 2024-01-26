package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"red": "#ff0000", "blue": "#0000ff", "green": "#00ff00"}
	// START OMIT
	fmt.Println(m)
	m["green"] = "verde" // HL
	fmt.Println(m)
	// END OMIT
}
