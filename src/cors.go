package main

import (
	"io"
	"log"
	"net/http"

	"github.com/jub0bs/cors"
)

func handleHello(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"hello":"world"}`)
}

// START OMIT
func main() {
	// create a custom ServeMux and register your handler to pattern GET /hello
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handleHello)

	// instantiate a CORS middleware whose config suits your needs
	corsMw, err := cors.NewMiddleware(cors.Config{ // HL
		Origins: []string{"*"}, // HL
	}) // HL
	if err != nil {
		log.Fatal(err)
	}

	// apply the CORS middleware
	handler := corsMw.Wrap(mux) // HL

	// start the server on port 8080; make sure to pass your custom handler
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}

// END OMIT
