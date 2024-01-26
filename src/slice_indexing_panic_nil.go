package main

import "fmt"

func main() {
	roles := []string{"Viewer", "Editor", "Admin"}
	// START OMIT
	roles = nil
	fmt.Println(roles[0]) // panic // HL
	// END OMIT
}
