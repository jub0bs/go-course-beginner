package main

import "fmt"

func main() {
	// START OMIT
	roles := []string{ // HL
		"Viewer",
		"Editor",
		"Admin", // this final comma is mandatory // HL
	}
	fmt.Printf("%T\n", roles)
	fmt.Println(roles)
	// END OMIT
}
