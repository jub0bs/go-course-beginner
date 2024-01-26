package main

import (
	"io"
	"log"
	"net/http"
)

// START OMIT
func main() {
	http.HandleFunc("GET /hello", handleHello)   // HL
	log.Fatal(http.ListenAndServe(":8080", nil)) // HL
}

func handleHello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello, world!")
}

// END OMIT
