package main

import (
	"fmt"
)

// START OMIT
func unhex(c byte) byte {
	if 'A' <= c && c <= 'F' || 'a' <= c && c <= 'f' {
		return c&0xDF - 'A' + 10
	} else if '0' <= c && c <= '9' {
		return c - '0'
	} else {
		return 0
	}
}

func main() {
	fmt.Println(unhex('F'))
}

// END OMIT
