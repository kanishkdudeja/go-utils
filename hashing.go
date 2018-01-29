package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword function generates the hash of the password string
// using the standard bcrypt library
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash functions checks if the supplied password when hashes equals the hash string suplied
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
