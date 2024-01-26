package main

import "fmt"

func main() {
	// START OMIT
	var count int
	go func() {
		count++ // equivalent to count = count + 1
	}()
	if count == 0 {
		fmt.Println(count)
	}
	// END OMIT
}
