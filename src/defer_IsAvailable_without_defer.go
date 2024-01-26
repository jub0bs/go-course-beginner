package main

import (
	"errors"
	"net/http"
)

func main() {
}

// START OMIT
func IsAvailable(username string) (bool, error) {
	resp, err := http.Get("https://github.com/" + username)
	if err != nil {
		return false, err
	}
	// omitted: use the response body // HL
	switch resp.StatusCode {
	case http.StatusNotFound:
		resp.Body.Close() // HL
		return true, nil
	case http.StatusOK:
		resp.Body.Close() // HL
		return false, nil
	default:
		resp.Body.Close() // HL
		return false, errors.New("unknown availability")
	}
}

// END OMIT
