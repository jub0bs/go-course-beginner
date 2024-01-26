package main

import (
	"io"
	"log"
	"net/http"

	"github.com/jub0bs/fcors"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"hello":"world"}`)
}

// START OMIT
func main() {
	// register the hello handler for the /hello path
	http.HandleFunc("/hello", hello)

	// instantiate a CORS middleware whose config suits your needs
	cors, err := fcors.AllowAccess( // HL
		fcors.FromAnyOrigin(), // HL
	) // HL
	if err != nil {
		log.Fatal(err)
	}

	// apply the CORS middleware to the default ServeMux
	mux := cors(http.DefaultServeMux) // HL

	// start the server on port 8080; make sure to path the modified ServeMux
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

// END OMIT
