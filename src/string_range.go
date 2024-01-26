package main

import (
	"fmt"
	"unicode/utf8"
)

// START OMIT
func main() {
	s := "Hello, 世界"
	for i, r := range s {
		const tmpl = "rune %q starts at index %d and is %d-byte long\n"
		fmt.Printf(tmpl, r, i, utf8.RuneLen(r))
	}
}

// END OMIT
