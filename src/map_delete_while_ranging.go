package main

import "fmt"

// START OMIT
func main() {
	m := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	for k := range m {
		if k == "foo" { // HL
			delete(m, "bar") // HL
		}
		if k == "bar" { // HL
			delete(m, "foo") // HL
		}
	}
	fmt.Println(m)
}

// END OMIT
