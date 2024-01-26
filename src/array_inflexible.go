package main

import "fmt"

// START OMIT
func concat3(a [3]string) string { // HL
	return a[0] + a[1] + a[2]
}

func main() {
	three := [3]string{"foo", "bar", "baz"}
	fmt.Println(concat3(three))
	four := [4]string{"foo", "bar", "baz", "qux"}
	fmt.Println(concat3(four)) // compilation error // HL
}

// END OMIT
