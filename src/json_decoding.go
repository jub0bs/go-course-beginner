package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`     // HL
	IsAdmin  bool   `json:"is_admin"` // HL
	verified bool
}

// START OMIT
func createUser(w http.ResponseWriter, r *http.Request) {
	// omitted: check that the request's content type is application/json
	var user User                             // HL
	dec := json.NewDecoder(r.Body)            // HL
	if err := dec.Decode(&user); err != nil { // HL
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// omitted: save user to the database
	w.WriteHeader(http.StatusCreated)
}

// END OMIT

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", createAdmin)
	if err := http.ListenAndServe(":8080", mux); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
