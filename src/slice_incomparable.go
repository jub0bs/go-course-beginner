package main

import "fmt"

func main() {
	// START OMIT
	userIds := []int{42, 56, 70}
	// fmt.Println(userIds == userIds) // compilation error // HL
	fmt.Println(userIds == nil) // HL

	roles := []string{"Viewer", "Editor", "Admin"}
	// fmt.Println(roles == roles) // compilation error // HL
	fmt.Println(roles == nil) // HL
	// END OMIT
}
