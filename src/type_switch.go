package main

import "fmt"

func main() {
	// START OMIT
	var v any = 42
	switch v := v.(type) { // HL
	case bool:
		fmt.Printf("bool %t\n", v)
	case int:
		fmt.Printf("int %d\n", v)
	default:
		fmt.Printf("unexpected type %T\n", v)
	}
	// END OMIT
}
