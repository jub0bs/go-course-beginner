package main

import "fmt"

var tmpl = "Hello, %s" // HL

func main() {
	name := "Julien" // HL
	fmt.Println(fmt.Sprintf(tmpl, name))
}
