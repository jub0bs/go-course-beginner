package main

import (
	"fmt"
)

// START OMIT
func unhex(c byte) byte {
	switch { // HL
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	default:
		return 0
	}
}

func main() {
	fmt.Println(unhex('e'))
}

// END OMIT
