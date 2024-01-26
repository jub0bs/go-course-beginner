package main

import (
	"fmt"
	"slices"
)

// START OMIT
type AuthMiddleware struct { // once instantiated, should ideally be immutable
	roles []string
}

func NewAuthMiddleware(roles ...string) *AuthMiddleware {
	return &AuthMiddleware{roles: roles} // lack of defensive copying // HL
}

func main() {
	roles := []string{"viewer", "editor", "admin"}
	mw := NewAuthMiddleware(roles...) // roles and mw.roles are now aliases    😕 // HL
	roles[0] = "superadmin"           // mutates the first element of mw.roles 😬 // HL
	fmt.Printf("%+v\n", mw)
}

// END OMIT
var _ = slices.Clone[[]int] // so as to import slices package without actually using it (yet)
