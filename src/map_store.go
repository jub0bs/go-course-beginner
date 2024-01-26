package main

import (
	"fmt"
)

func main() {
	// START OMIT
	m := map[string]string{"red": "#ff0000", "blue": "#0000ff"}
	m["green"] = "#00ff00" // HL
	fmt.Println(m)
	// END OMIT
}
