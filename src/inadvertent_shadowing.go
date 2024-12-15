package main

import (
	"errors"
	"fmt"
	"strconv"
)

// START OMIT
func AllParsableAsPositiveInts(inputs ...string) error {
	var err error // HL
	for _, str := range inputs {
		n, err := strconv.Atoi(str) // this err now shadows the one declared in the outer scope // HL
		if err != nil {
			break
		}
		if n <= 0 {
			err = errors.New("nonpositive integer") // doesn't update the outer err variable! // HL
			break
		}
	}
	return err
}

func main() {
	fmt.Println(AllParsableAsPositiveInts("49.3", "-1")) // nil 🤔 // HL
}

// END OMIT
