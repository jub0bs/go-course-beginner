package main

import (
	"fmt"
	"strings"
)

func main() {
	// START OMIT
	fmt.Println(strings.ToUpper == strings.ToLower) // compilation error!
	// END OMIT
}
