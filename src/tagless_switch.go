package main

import (
	"fmt"
)

// START OMIT
func unhex(c byte) byte {
	switch { // HL
	case 'A' <= c && c <= 'F' || 'a' <= c && c <= 'f':
		return c&0xDF - 'A' + 10
	case '0' <= c && c <= '9':
		return c - '0'
	default:
		return 0
	}
}

func main() {
	fmt.Println(unhex('F'))
}

// END OMIT
