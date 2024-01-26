package main

import "fmt"

// START OMIT
var i = -42

func main() {
	s := "Hello, 世界"
	fmt.Println("length:", len(s))
	fmt.Print("bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()

	// fmt.Println(s[i]) // panic!
	// fmt.Println(s[len(s)]) // panic!

	// s[0] = 'h' // compilation error
	// fmt.Println(&s[i]) // compilation error
}

// END OMIT
