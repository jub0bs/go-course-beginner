package main

import "fmt"

var n = 16

func main() {
	n := 8 // shadows package-scoped variable n
	fmt.Println(n)
}
