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
	defer func(e error) { // HL
		err = fmt.Errorf("github: %w", e) // HL
	}(err) // HL
	// real implementation omitted
	switch {
	// other cases omitted
	default:
		return false, errors.New("unknown availability")
	}
}

// END OMIT
