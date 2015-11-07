// Package data provides utilities for interacting with database data
// throughout Write.as.
package data

import (
	"bytes"
	"crypto/rand"
	"strings"
	"testing"
)

func TestEncDec(t *testing.T) {
	// Generate a random key with a valid length
	k := make([]byte, keyLen)
	_, err := rand.Read(k)
	if err != nil {
		t.Fatal(err)
	}

	runEncDec(t, k, "this is my secret message™. 😄", nil)
	runEncDec(t, k, "mygreatemailaddress@gmail.com", nil)
}

func TestAuthentication(t *testing.T) {
	// Generate a random key with a valid length
	k := make([]byte, keyLen)
	_, err := rand.Read(k)
	if err != nil {
		t.Fatal(err)
	}

	runEncDec(t, k, "mygreatemailaddress@gmail.com", func(c []byte) []byte {
		c[0] = 'a'
		t.Logf("Modified:   %s\n", c)
		return c
	})
}

func runEncDec(t *testing.T, k []byte, plaintext string, transform func([]byte) []byte) {
	t.Logf("Plaintext:  %s\n", plaintext)

	// Encrypt the data
	ciphertext, err := Encrypt(k, plaintext)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Ciphertext: %s\n", ciphertext)

	if transform != nil {
		ciphertext = transform(ciphertext)
	}

	// Decrypt the data
	decryptedText, err := Decrypt(k, ciphertext)
	if err != nil {
		if transform != nil && strings.Contains(err.Error(), "message authentication failed") {
			// We modified the ciphertext; make sure we're getting the right error
			t.Logf("%v\n", err)
			return
		}
		t.Fatal(err)
	}

	t.Logf("Decrypted:  %s\n", string(decryptedText))

	if !bytes.Equal([]byte(plaintext), decryptedText) {
		t.Errorf("Plaintext mismatch: got %x vs %x", plaintext, decryptedText)
	}
}
