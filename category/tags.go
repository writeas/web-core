package category

import (
	"strings"
	"unicode"
)

// titleFromHashtag generates an all-lowercase title, with spaces inserted based on initial capitalization -- e.g.
// "MyWordyTag" becomes "my wordy tag".
func titleFromHashtag(hashtag string) string {
	var t strings.Builder
	var prev rune
	for i, c := range hashtag {
		if unicode.IsUpper(c) {
			if i > 0 && !unicode.IsUpper(prev) {
				// Insert space if previous rune wasn't also uppercase (e.g. an abbreviation)
				t.WriteRune(' ')
			}
			t.WriteRune(unicode.ToLower(c))
		} else {
			t.WriteRune(c)
		}
		prev = c
	}
	return t.String()
}

// HashtagFromTitle generates a valid single-word, camelCase hashtag from a title (which might include spaces,
// punctuation, etc.).
func HashtagFromTitle(title string) string {
	var t strings.Builder
	var prev rune
	for _, c := range title {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			prev = c
			continue
		}
		if unicode.IsSpace(prev) {
			// Uppercase next word
			t.WriteRune(unicode.ToUpper(c))
		} else {
			t.WriteRune(c)
		}
		prev = c
	}
	return t.String()
}
