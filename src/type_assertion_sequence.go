package main

import "fmt"

func main() {
	// START OMIT
	var v any = 42
	if b, ok := v.(bool); ok {
		fmt.Printf("bool %t\n", b)
	} else if i, ok := v.(int); ok {
		fmt.Printf("int %d\n", i)
	} else {
		fmt.Printf("unexpected type %T\n", v)
	}
	// END OMIT
}
