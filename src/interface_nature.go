package main

import "fmt"

func main() {
	// START OMIT
	var i any = 42 // dynamic type: int;    dynamic value: 42
	i = "foo"      // dynamic type: string; dynamic value: "foo"
	i = false      // dynamic type: bool;   dynamic value: false
	fmt.Println(i)
	// END OMIT
}
