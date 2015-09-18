// Package passgen generates random, unmemorable passwords.
//
// Example usage:
//
//	p := passgen.New() // p is "6NX(W`GD]4:Tqk};Y@A-"
//
// Logic originally from dchest's uniuri library:
// https://github.com/dchest/uniuri
//
// Functions read from crypto/rand random source, and panic if they fail to
// read from it.
package passgen

import "crypto/rand"

// DefLen is the default password length returned.
const DefLen = 20

// DefChars is the default set of characters used in the password.
var DefChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()- _=+,.?/:;{}[]`~")

// New returns a random password of the default length with the default set of
// characters.
func New() string {
	return NewLenChars(DefLen, DefChars)
}

// NewLen returns a random password of the given length with the default set
// of characters.
func NewLen(length int) string {
	return NewLenChars(length, DefChars)
}

// NewLenChars returns a random password of the given length with the given
// set of characters.
func NewLenChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("passgen: error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}
