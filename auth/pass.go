package auth

import "golang.org/x/crypto/bcrypt"

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

func HashPass(password []byte) ([]byte, error) {
	// Clear memory where plaintext password was stored.
	// http://stackoverflow.com/questions/18545676/golang-app-engine-securely-hashing-a-users-password#comment36585613_19828153
	defer clear(password)
	// Return hash
	return bcrypt.GenerateFromPassword(password, 12)
}

func Authenticated(hash, pass []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, pass) == nil
}
