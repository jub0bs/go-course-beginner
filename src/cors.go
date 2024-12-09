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
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handleHello)

	// instantiate a CORS middleware whose configuration suits your needs
	corsMw, err := cors.NewMiddleware(cors.Config{ // HL
		Origins: []string{"https://example.com"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// apply your CORS middleware to your HTTP request multiplexer
	handler := corsMw.Wrap(mux) // HL

	// pass the result to ListenAndServe
	if err := http.ListenAndServe(":8080", handler); err != http.ErrServerClosed { // HL
		log.Fatal(err)
	}
}

// END OMIT
