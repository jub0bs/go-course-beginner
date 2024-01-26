package main

import "fmt"

func main() {
	roles := []string{"Viewer", "Editor", "Admin"}
	// START OMIT
	fmt.Println(roles[len(roles)]) // panic // HL
	// END OMIT
}
