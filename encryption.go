package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// EncryptByteArray function encrypts a byte array using a key and returns the encrypted byte array
// If the encryption fails, it returns an error
// The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func EncryptByteArray(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// DecryptByteArray function decrypts the ciphertext byte array using the key provided
// If the decryption is successful, it returns the plain text byte array. If the decryption fails, it returns an error
// The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func DecryptByteArray(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

// EncryptString takes plain text and key as string arguments
// and returns the encrypted string
// It calls the EncryptByteArray for the core encryption logic
// The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func EncryptString(plainText string, key string) (string, error) {
	encryptedByteArray, err := EncryptByteArray([]byte(plainText), []byte(key))

	if err != nil {
		return "", err
	}

	return string(encryptedByteArray), nil
}

// DecryptString takes cipher text and key as string arguments
// and returns the decrypted string
// It calls the DecryptByteArray for the core encryption logic
// The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func DecryptString(cipherText string, key string) (string, error) {
	decryptedByteArray, err := DecryptByteArray([]byte(cipherText), []byte(key))

	if err != nil {
		return "", err
	}

	return string(decryptedByteArray), nil
}
