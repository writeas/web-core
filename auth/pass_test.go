package auth

import "testing"

const pass = "password"

var hash []byte

func TestHash(t *testing.T) {
	var err error
	hash, err = HashPass([]byte(pass))
	if err != nil {
		t.Error("Password hash failed.")
	}
}

func TestAuth(t *testing.T) {
	if !Authenticated(hash, []byte(pass)) {
		t.Error("Didn't authenticate.")
	}
}
