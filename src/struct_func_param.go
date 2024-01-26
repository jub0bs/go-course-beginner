package main

import "fmt"

// START OMIT
type User struct {
	firstName string
	lastName  string
	isAdmin   bool
}

func Promote(u User) { // HL
	u.isAdmin = true
}

func main() {
	julien := User{
		lastName:  "Cretel",
		firstName: "Julien",
	}
	Promote(julien)
	fmt.Println(julien.isAdmin) // false
}

// END OMIT
