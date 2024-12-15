package main

import (
	"fmt"
	"strings"
)

func main() {
	// START OMIT
	// (Function say remains as shown earlier.)
	say("Yo! ", func(what string) string { return strings.Repeat(what, 3) })
	// END OMIT
}

func say(what string, how func(string) string) {
	fmt.Println(how(what))
}
