package tags

import (
	"github.com/kylemcc/twitter-text-go/extract"
)

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
