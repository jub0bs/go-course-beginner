package main

import "fmt"

func main() {
	// START OMIT
	i := 42
	j := i
	fmt.Println(&i == &j) // false: i and j are distinct variables
	// END OMIT
}
