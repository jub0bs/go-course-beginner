package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

// START OMIT
func main() {
	const input = "foo\x00bar"
	fmt.Println(Base64Encode(input)) // want "Zm9vAGJhcg=="
}

func Base64Encode(s string) string {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	defer enc.Close() // HL
	io.WriteString(enc, s)
	return buf.String()
}

// END OMIT
