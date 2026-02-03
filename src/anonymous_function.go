package main

import (
	"fmt"
	"strings"
)

	// START OMIT
func main() {
	say("Yo! ", func(what string) string { return strings.Repeat(what, 3) }) // HL
}

func say(what string, how func(string) string) {
	fmt.Println(how(what))
}
	// END OMIT
