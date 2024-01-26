package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"red": "#ff0000", "blue": "#0000ff", "empty": ""}
	// START OMIT
	v, ok = m["green"] // HL
	fmt.Println(v, ok)
	v, ok := m["empty"] // HL
	fmt.Println(v, ok)
	// END OMIT
}
