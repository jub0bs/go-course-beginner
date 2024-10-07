package main

import (
	"fmt"
	"mime"
)

func main() {
	// START OMIT
	mediatype, _, err := mime.ParseMediaType(`image/png; ¯\_(ツ)_/¯`)
	fmt.Printf("%q, %q\n", mediatype, err)
	// END OMIT
}
