package main

import "fmt"

// START OMIT
type User struct {
	name    string
	isAdmin bool
	// friends []User // this field's type is not comparable // HL
}

func main() {
	alice := User{name: "Alice", isAdmin: true}
	bob := User{name: "Bob"}
	fmt.Println(alice == bob)
}

// END OMIT
