package text

import (
	"github.com/kylemcc/twitter-text-go/extract"
	"net/url"
	"regexp"
)

var imageURLRegex = regexp.MustCompile(`(?i)[^ ]+\.(gif|png|jpg|jpeg|image)$`)

func ExtractImages(content string) []string {
	matches := extract.ExtractUrls(content)
	urls := map[string]bool{}
	for i := range matches {
		uRaw := matches[i].Text
		// Parse the extracted text so we can examine the path
		u, err := url.Parse(uRaw)
		if err != nil {
			continue
		}
		// Ensure the path looks like it leads to an image file
		if !imageURLRegex.MatchString(u.Path) {
			continue
		}
		urls[uRaw] = true
	}

	resURLs := make([]string, 0)
	for k := range urls {
		resURLs = append(resURLs, k)
	}
	return resURLs
}
