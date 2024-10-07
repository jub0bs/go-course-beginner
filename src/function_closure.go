package main

import (
	"fmt"
)

// START OMIT
func main() {
	f := falseAfterN(4)
	fmt.Println(f(), f(), f(), f(), f(), f())
}

func falseAfterN(n uint) func() bool {
	var count uint // captured by the anonymous function below // HL
	return func() bool {
		if count < n { // HL
			count++ // HL
			return true
		}
		return false
	}
}

// END OMIT
