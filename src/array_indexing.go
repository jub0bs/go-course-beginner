package main

import "fmt"

// START OMIT
var i = 3

func main() {
	roles := [...]string{"Viewer", "Editor", "Admin"}
	fmt.Println(&roles[0]) // HL
	fmt.Println(&roles[1])
	fmt.Println(&roles[2])
	// fmt.Println(roles[i]) // panic // HL
	// fmt.Println(roles[3]) // compilation error // HL
}

// END OMIT
