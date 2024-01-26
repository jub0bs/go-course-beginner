package main

import (
	"errors"
	"net/http"
	"net/url"
)

func main() {
}

// START OMIT
func IsAvailable(username string) (bool, error) {
	addr, err := url.JoinPath("https://github.com", url.PathEscape(username))
	if err != nil {
		return false, err
	}
	resp, err := http.Get(addr)
	if err != nil {
		return false, err
	}
	switch resp.StatusCode {
	case http.StatusNotFound:
		resp.Body.Close() // 😬 // HL
		return true, nil
	case http.StatusOK:
		resp.Body.Close() // 😬 // HL
		return false, nil
	default:
		resp.Body.Close() // 😬 // HL
		return false, errors.New("unknown availability")
	}
}

// END OMIT