package id

import "testing"

func TestGenSafeUniqueSlug(t *testing.T) {
	slug := "slug"
	r := map[string]bool{}

	for i := 0; i < 1000; i++ {
		s := GenSafeUniqueSlug(slug)
		if s == slug {
			t.Errorf("Got same slug as inputted!")
		}
		if _, ok := r[s]; ok {
			t.Errorf("#%d: slug %s was already generated in testing.", i, s)
		}
		r[s] = true
	}
}
