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
	// create a custom HTTP request multiplexer and register your handler to pattern GET /hello
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handleHello)

	// instantiate a CORS middleware whose config suits your needs
	corsMw, err := cors.NewMiddleware(cors.Config{
		Origins: []string{"https://example.com"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// apply your CORS middleware to your HTTP request multiplexer
	handler := corsMw.Wrap(mux)

	// start the server on port 8080; make sure to use your custom handler
	if err := http.ListenAndServe(":8080", handler); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// END OMIT
