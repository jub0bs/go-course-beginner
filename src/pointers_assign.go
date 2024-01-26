package main

import "fmt"

func main() {
	// START OMIT
	i := 42
	p := &i
	*p++
	fmt.Println(i)
	// END OMIT
}
