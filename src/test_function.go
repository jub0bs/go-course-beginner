package github_test // external test package: black-box testing

import (
	"testing"
	// omitted: import declaration of your github package
)

func TestUsernameContainsTwoConsecutiveHyphens(t *testing.T) { // HL
	const username = "jub0bs--on-GitHub"
	const want = false
	got := github.IsValid(username)
	if got != want {
		t.Errorf("github.IsValid(%q): got %t; want %t", username, got, want) // HL
	}
}
