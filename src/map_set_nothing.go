package main

import (
	"fmt"
)

func main() {
	// START OMIT
	set := map[string]struct{}{ // this map's keys represent the set's elements
		"red":   {},
		"blue":  {},
		"green": {},
	}

	s := "yellow"
	if _, ok := set[s]; ok { // checking for membership
		fmt.Printf("%q is not in the map.\n", s)
	}

	set["white"] = struct{}{} // adding an element

	delete(set, "red") // deleting an element

	fmt.Println(set)
	// END OMIT
}
