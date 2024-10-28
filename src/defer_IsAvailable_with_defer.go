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
	defer resp.Body.Close() // 😌 // HL
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
