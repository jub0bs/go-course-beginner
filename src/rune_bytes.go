package main

// START OMIT
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const tmpl = "%q is made up of %d bytes\n"
	const ascii = 'a'
	const chinese = '世'
	fmt.Printf(tmpl, ascii, utf8.RuneLen(ascii))
	fmt.Printf(tmpl, chinese, utf8.RuneLen(chinese))
}
// END OMIT

