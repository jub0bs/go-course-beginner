package main

import "fmt"

var tmpl = "Hello, %s" // HL

func main() {
	name := "Julien" // HL
	fmt.Printf("type: %T\n", name)
	fmt.Println(fmt.Sprintf(tmpl, name))
}
