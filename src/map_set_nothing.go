package main

import (
	"fmt"
)

func main() {
	// START OMIT
	var present struct{}
	m := map[string]struct{}{"red": present, "blue": present, "green": present}
	const k = "yellow"

	// checking for membership
	if _, ok := m[k]; !ok {
		fmt.Printf("%q is not in the map.\n", k)
	}

	// adding an element
	m["white"] = present

	// deleting an element
	delete(m, "red")

	fmt.Println(m)
	// END OMIT
}
