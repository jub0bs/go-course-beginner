package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	names := []string{"foo", "bar", "baz"}
	s := Concat(names...) // HL
	fmt.Println(s)
}

func Concat(rest ...string) string {
	return strings.Join(rest, "-")
}

// END OMIT
