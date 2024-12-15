package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	say("What's for dinner?", strings.ToUpper)
}

func say(what string, how func(string) string) {
	fmt.Println(how(what))
}

// END OMIT
