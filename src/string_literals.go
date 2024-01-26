package main

import "fmt"

// START OMIT
func main() {
	var s string
	fmt.Printf("%q\n", s)
	fmt.Println(s == "")

	msg := "Hello, World!"
	fmt.Println(msg)

	invalidUTF8 := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(invalidUTF8)

	raw := `This is a multiline string literal.
Escapes like \xbd don't work here.`
	fmt.Println(raw)
}
// END OMIT

