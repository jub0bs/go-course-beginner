package main

import (
	"fmt"
)

func main() {
	// START OMIT
	m := map[string]string{"red": "#ff0000", "blue": "#0000ff", "green": "#00ff00"}
	fmt.Println(m)

	delete(m, "red") // HL
	fmt.Println(m)

	delete(m, "red")
	fmt.Println(m)
	// END OMIT
}
