package main

import "fmt"

func main() {
	// START OMIT
	roles := [...]string{"Viewer", "Editor", "Admin"}
	fmt.Println(roles == roles) // legal because type string is comparable // HL

	var funcs [0]func()
	fmt.Println(funcs)
	// fmt.Println(funcs == funcs) // illegal because type func() is not comparable // HL
	// END OMIT
}
