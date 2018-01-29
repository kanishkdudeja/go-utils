package utils

import (
	"regexp"
)

// IsEmailValid uses a standard Regex expresssion to check if an email address is valid or not
func IsEmailValid(email string) bool {
	Re := regexp.MustCompile(`^([a-zA-Z0-9+_\-\.]+)@([a-zA-Z0-9+_\-\.]+)\.([a-zA-Z]{2,5})$`)

	return Re.MatchString(email)
}
