package main

import (
	"io"
	"log"
	"net/http"
)

// START OMIT
func main() {
	http.HandleFunc("GET /hello", handleHello) // HL
	if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleHello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello, world!")
}

// END OMIT
