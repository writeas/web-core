package id

import (
	"crypto/rand"
	"fmt"
)

// GenerateRandomString creates a random string of characters of the given
// length from the given dictionary of possible characters.
//
// This example generates a hexadecimal string 6 characters long:
//     GenerateRandomString("0123456789abcdef", 6)
func GenerateRandomString(dictionary string, l int) string {
	var bytes = make([]byte, l)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

// GenSafeUniqueSlug generatees a reasonably unique random slug from the given
// original slug. It's "safe" because it uses 0-9 b-z excluding vowels.
func GenSafeUniqueSlug(slug string) string {
	return fmt.Sprintf("%s-%s", slug, GenerateRandomString("0123456789bcdfghjklmnpqrstvwxyz", 4))
}

// Generate62RandomString creates a random string with the given length
// consisting of characters in [A-Za-z0-9].
func Generate62RandomString(l int) string {
	return GenerateRandomString("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", l)
}

// GenerateFriendlyRandomString creates a random string of characters with the
// given length consisting of characters in [a-z0-9].
func GenerateFriendlyRandomString(l int) string {
	return GenerateRandomString("0123456789abcdefghijklmnopqrstuvwxyz", l)
}
