package main

import (
	"fmt"
)

// START OMIT
func main() {
	fmt.Println(greet("Hello, %s!", "Julien"))
}

func greet(tmpl, name string) string { // HL
	return fmt.Sprintf(tmpl, name)
}

// END OMIT
