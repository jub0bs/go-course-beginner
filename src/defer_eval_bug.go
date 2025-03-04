package main

import (
	"errors"
	"fmt"
	"log"
)

// START OMIT
func main() {
	avail, err := IsAvailable("jub0bs")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(avail)
}

func IsAvailable(username string) (avail bool, err error) {
	defer func(err1 error) { // HL
		if err1 != nil {
			err = fmt.Errorf("github: %w", err1) // HL
		}
	}(err) // HL
	// real implementation omitted
	return false, errors.New("unknown availability")
}

// END OMIT
