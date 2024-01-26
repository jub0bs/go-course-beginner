package main

import "fmt"

// START OMIT
type User struct {
	firstName string
	lastName  string
	isAdmin   bool
}

func main() {
	var u User
	fmt.Printf("%+v\n", u)
}

// END OMIT
