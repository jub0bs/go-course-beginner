package main

import "fmt"

func main() {
	roles := [...]string{"Viewer", "Editor", "Admin"}
	// START OMIT
	for _, role := range roles { // HL
		fmt.Println(role)
	}

	for i := range roles { // HL
		fmt.Println(&roles[i])
	}
	// END OMIT
}
