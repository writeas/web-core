// Package tags supports operations around hashtags in plain text content
package tags

import (
	"github.com/kylemcc/twitter-text-go/extract"
)

// Extract finds all hashtags in the given string and returns a de-duplicated
// list of them.
func Extract(body string) []string {
	matches := extract.ExtractHashtags(body)
	tags := map[string]bool{}
	for i := range matches {
		// Second value (whether or not there's a hashtag) ignored here, since
		// we're only extracting hashtags.
		ht, _ := matches[i].Hashtag()
		tags[ht] = true
	}

	resTags := make([]string, 0)
	for k := range tags {
		resTags = append(resTags, k)
	}
	return resTags
}
