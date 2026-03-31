package main

import "fmt"

// START OMIT
type ValidationError struct{} // fields omitted for brevity

func (*ValidationError) Error() string { return "not a palindrome" }

func ValidateBytePalindrome(s string) error { // HL
	var err *ValidationError // HL
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			err = &ValidationError{} // HL
			break
		}
	}
	return err // HL
}

func main() {
	fmt.Println(ValidateBytePalindrome("level") == nil) // true or false? // HL
}

// END OMIT
