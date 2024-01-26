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
	defer resp.Body.Close() // HL
	// omitted: use the response body // HL
	switch resp.StatusCode {
	case http.StatusNotFound:
		return true, nil
	case http.StatusOK:
		return false, nil
	default:
		return false, errors.New("unknown availability")
	}
}

// END OMIT
