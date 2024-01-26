package main

import (
	"fmt"
)

func main() {
	// START OMIT
	m := map[string]string{"red": "#ff0000", "blue": "#0000ff", "empty": ""}
	v := m["red"] // HL
	fmt.Println(v)
	// END OMIT
}
