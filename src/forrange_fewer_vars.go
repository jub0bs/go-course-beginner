package main

import "fmt"

func main() {
	// START OMIT
	var runeCount uint
	for range "Hello, 世界" {
		runeCount++
	}
	fmt.Println(runeCount)
	// END OMIT
}
