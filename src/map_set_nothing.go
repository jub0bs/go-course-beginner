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

	s := "blue"
	if _, ok := set[s]; ok { // checking for membership
		fmt.Printf("%q is in the set.\n", s)
	}

	set["white"] = struct{}{} // adding an element

	delete(set, "red") // deleting an element

	fmt.Println(set)
	// END OMIT
}
