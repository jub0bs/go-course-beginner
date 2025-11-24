package main

import (
	"errors"
	"net/http"
)

func main() {
}

// START OMIT
func IsAvailable(username string) (bool, error) {
	// ...
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	switch res.StatusCode {
	case http.StatusNotFound:
		res.Body.Close() // 😬 // HL
		return true, nil
	case http.StatusOK:
		res.Body.Close() // 😬 // HL
		return false, nil
	default:
		res.Body.Close() // 😬 // HL
		return false, errors.New("unknown availability")
	}
}

// END OMIT
