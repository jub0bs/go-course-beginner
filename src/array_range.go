package main

import "fmt"

func main() {
	// START OMIT
	roles := [...]string{"Viewer", "Editor", "Admin"}
	for i, role := range roles { // HL
		fmt.Printf("%d: %q, %t\n", i, role, &role == &roles[i])
	}
	// END OMIT
}
