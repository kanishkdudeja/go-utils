package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"

	"github.com/satori/go.uuid"
)

func byte2string(in [16]byte) []byte {
	return in[:16]
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)

	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

// GenerateARandomString returns a 32 digit long random string
func GenerateARandomString() (string, error) {
	randomString, err := GenerateRandomString(32)

	if err != nil {
		return "", err
	}

	data := []byte(randomString)
	return hex.EncodeToString(byte2string(md5.Sum(data))), nil
}

// GetNewUUID returns a V4 UUID
// as per the https://tools.ietf.org/html/rfc4122#section-4.4 spec
func GetNewUUID() (string, error) {
	newUUID, err := uuid.NewV4()

	if err != nil {
		return "", err
	}

	return newUUID.String(), nil
}
