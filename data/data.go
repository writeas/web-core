// Package data provides utilities for interacting with database data
// throughout Write.as.
package data

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
)

// Encryption parameters
const (
	keyLen    = 32
	delimiter = '%'
)

// Encrypt AES-encrypts given text with the given key k.
// This is used for encrypting sensitive information in the database, such as
// oAuth tokens and email addresses.
func Encrypt(k []byte, text string) ([]byte, error) {
	// Validate parameters
	if len(k) != keyLen {
		return nil, errors.New(fmt.Sprintf("Invalid key length (must be %d bytes).", keyLen))
	}

	// Encrypt plaintext with AES-GCM
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Generate nonce
	ns := gcm.NonceSize()
	nonce := make([]byte, ns)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(text), nil)

	// Build text output in the format:
	// NonceCiphertext
	outtext := append(nonce, ciphertext...)

	return outtext, nil
}

// Decrypt decrypts the given ciphertext with the given key k.
func Decrypt(k, ciphertext []byte) ([]byte, error) {
	// Decrypt ciphertext
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ns := gcm.NonceSize()

	// Validate data
	if len(ciphertext) < ns {
		return nil, errors.New("Ciphertext is too short")
	}

	nonce := ciphertext[:ns]
	ciphertext = ciphertext[ns:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
