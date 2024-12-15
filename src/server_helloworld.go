package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT
func main() {
	mux := http.NewServeMux()                                                  // HL
	mux.HandleFunc("GET /hello", handleHello)                                  // HL
	if err := http.ListenAndServe(":8080", mux); err != http.ErrServerClosed { // HL
		log.Fatal(err)
	}
}

func handleHello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

// END OMIT
