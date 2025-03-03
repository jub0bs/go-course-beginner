package main

import (
	"fmt"
	"log"
	"regexp"
)

// START OMIT
var re *regexp.Regexp // pointer to a regexp.Regexp // HL

func main() {
	re, err := regexp.Compile("^[0-9a-z]{,8}$") // HL
	if err != nil {
		log.Fatal(err)
	}
	re.Longest()
	fmt.Println(IsValid("jub0bs"))
}

func IsValid(username string) bool {
	return re.MatchString(username) // panic! // HL
}

// END OMIT
