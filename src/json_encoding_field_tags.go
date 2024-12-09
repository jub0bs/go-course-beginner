package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// START OMIT
type User struct {
	Name     string `json:"name"`     // HL
	IsAdmin  bool   `json:"is_admin"` // HL
	verified bool
}

// END OMIT

func getUser(w http.ResponseWriter, r *http.Request) {
	// omitted: retrieve the requested user from the database
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)                // HL
	if err := enc.Encode(user); err != nil { // HL
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", getUser)
	if err := http.ListenAndServe(":8080", mux); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
