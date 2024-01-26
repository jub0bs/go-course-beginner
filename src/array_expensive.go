package main

import "fmt"

// START OMIT
func main() {
	var a [1000]int
	fmt.Println(&a[0])
	printAddrOfFirstElement(a)
}

func printAddrOfFirstElement(a [1000]int) {
	fmt.Println(&a[0])
}
// END OMIT
