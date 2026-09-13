package main

import (
	"encoding/json/v2"
	"log"
	"net/http"
)

// START OMIT
type User struct {
	Name     string
	IsAdmin  bool
	verified bool
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var user User // omitted: retrieve the requested user from the database
	w.Header().Set("Content-Type", "application/json")
	if err := json.MarshalWrite(w, user); err != nil { // HL
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// END OMIT

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", getUser)
	if err := http.ListenAndServe(":8080", mux); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
