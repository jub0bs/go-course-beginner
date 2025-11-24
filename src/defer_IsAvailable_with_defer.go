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
	res, err := http.DefaultCLient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close() // 😌 // HL
	switch res.StatusCode {
	case http.StatusNotFound:
		return true, nil
	case http.StatusOK:
		return false, nil
	default:
		return false, errors.New("unknown availability")
	}
}

// END OMIT
