package id

import (
	"fmt"
	"github.com/writeas/nerds/store"
)

// GenSafeUniqueSlug generatees a reasonably unique random slug from the given
// original slug. It's "safe" because it uses 0-9 b-z excluding vowels.
func GenSafeUniqueSlug(slug string) string {
	return fmt.Sprintf("%s-%s", slug, store.GenerateRandomString("0123456789bcdfghjklmnpqrstvwxyz", 4))
}
