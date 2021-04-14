package posts

import (
	"fmt"
	stripmd "github.com/writeas/go-strip-markdown"
	"github.com/writeas/web-core/stringmanip"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	maxTitleLen     = 80
	assumedTitleLen = 80
)

var (
	titleElementReg = regexp.MustCompile("</?p>")
	urlReg          = regexp.MustCompile("https?://")
	imgReg          = regexp.MustCompile(`!\[([^]]+)\]\([^)]+\)`)
)

// ExtractTitle takes the given raw post text and returns a title, if explicitly
// provided, and a body.
func ExtractTitle(content string) (title string, body string) {
	if hashIndex := strings.Index(content, "# "); hashIndex == 0 {
		eol := strings.IndexRune(content, '\n')
		// First line should start with # and end with \n
		if eol != -1 {
			body = strings.TrimLeft(content[eol:], " \t\n\r")
			title = content[len("# "):eol]
			return
		}
	}
	body = content
	return
}

func FriendlyPostTitle(content, friendlyId string) string {
	content = StripHTMLWithoutEscaping(content)

	content = strings.TrimLeftFunc(stripmd.Strip(content), unicode.IsSpace)
	eol := strings.IndexRune(content, '\n')
	blankLine := strings.Index(content, "\n\n")
	if blankLine != -1 && blankLine <= eol && blankLine <= assumedTitleLen {
		return strings.TrimSpace(content[:blankLine])
	} else if eol == -1 && utf8.RuneCountInString(content) <= maxTitleLen {
		return content
	}

	title, truncd := TruncToWord(PostLede(content, true), maxTitleLen)
	if truncd {
		title += "..."
	}
	return title
}

// PostDescription generates a description based on the given post content,
// title, and post ID. This doesn't consider a V2 post field, `title` when
// choosing what to generate. In case a post has a title, this function will
// fail, and logic should instead be implemented to skip this when there's no
// title, like so:
//    var desc string
//    if title == "" {
//        desc = PostDescription(content, title, friendlyId)
//    } else {
//        desc = ShortPostDescription(content)
//    }
func PostDescription(content, title, friendlyId string) string {
	maxLen := 140

	if content == "" {
		content = "WriteFreely is a painless, simple, federated blogging platform."
	} else {
		fmtStr := "%s"
		truncation := 0
		if utf8.RuneCountInString(content) > maxLen {
			// Post is longer than the max description, so let's show a better description
			fmtStr = "%s..."
			truncation = 3
		}

		if title == friendlyId {
			// No specific title was found; simply truncate the post, starting at the beginning
			content = fmt.Sprintf(fmtStr, strings.Replace(stringmanip.Substring(content, 0, maxLen-truncation), "\n", " ", -1))
		} else {
			// There was a title, so return a real description
			blankLine := strings.Index(content, "\n\n")
			if blankLine < 0 {
				blankLine = 0
			}
			truncd := stringmanip.Substring(content, blankLine, blankLine+maxLen-truncation)
			contentNoNL := strings.Replace(truncd, "\n", " ", -1)
			content = strings.TrimSpace(fmt.Sprintf(fmtStr, contentNoNL))
		}
	}

	return content
}

func ShortPostDescription(content string) string {
	maxLen := 140
	fmtStr := "%s"
	truncation := 0
	if utf8.RuneCountInString(content) > maxLen {
		// Post is longer than the max description, so let's show a better description
		fmtStr = "%s..."
		truncation = 3
	}
	return strings.TrimSpace(fmt.Sprintf(fmtStr, strings.Replace(stringmanip.Substring(content, 0, maxLen-truncation), "\n", " ", -1)))
}

// TruncToWord truncates the given text to the provided limit.
func TruncToWord(s string, l int) (string, bool) {
	truncated := false
	c := []rune(s)
	if len(c) > l {
		truncated = true
		s = string(c[:l])
		spaceIdx := strings.LastIndexByte(s, ' ')
		if spaceIdx > -1 {
			s = s[:spaceIdx]
		}
	}
	return s, truncated
}

// PostLede attempts to extract the first thought of the given post, generally
// contained within the first line or sentence of text.
func PostLede(t string, includePunc bool) string {
	// Adjust where we truncate if we want to include punctuation
	iAdj := 0
	if includePunc {
		iAdj = 1
	}

	// Find lede within first line of text
	nl := strings.IndexRune(t, '\n')
	if nl > -1 {
		t = t[:nl]
	}

	// Strip certain HTML tags
	t = titleElementReg.ReplaceAllString(t, "")

	// Strip URL protocols
	t = urlReg.ReplaceAllString(t, "")

	// Strip image URL, leaving only alt text
	t = imgReg.ReplaceAllString(t, " $1 ")

	// Find lede within first sentence
	punc := strings.Index(t, ". ")
	if punc > -1 {
		t = t[:punc+iAdj]
	}
	punc = stringmanip.IndexRune(t, '。')
	if punc > -1 {
		c := []rune(t)
		t = string(c[:punc+iAdj])
	}

	return t
}
