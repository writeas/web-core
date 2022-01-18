package posts

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/writeas/saturday"
	"html"
	"regexp"
	"strings"
	"unicode"
)

var (
	blockReg    = regexp.MustCompile("<(ul|ol|blockquote)>\n")
	endBlockReg = regexp.MustCompile("</([a-z]+)>\n</(ul|ol|blockquote)>")

	markeddownReg = regexp.MustCompile("<p>(.+)</p>")
)

func ApplyMarkdown(data []byte) string {
	mdExtensions := 0 |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS
	htmlFlags := 0 |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_DASHES

	// Generate Markdown
	md := blackfriday.Markdown([]byte(data), blackfriday.HtmlRenderer(htmlFlags, "", ""), mdExtensions)
	// Strip out bad HTML
	policy := bluemonday.UGCPolicy()
	policy.AllowAttrs("class", "id").Globally()
	outHTML := string(policy.SanitizeBytes(md))
	// Strip newlines on certain block elements that render with them
	outHTML = blockReg.ReplaceAllString(outHTML, "<$1>")
	outHTML = endBlockReg.ReplaceAllString(outHTML, "</$1></$2>")

	return outHTML
}

func ApplyBasicMarkdown(data []byte) string {
	mdExtensions := 0 |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS
	htmlFlags := 0 |
		blackfriday.HTML_SKIP_HTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_DASHES

	// Generate Markdown
	md := blackfriday.Markdown([]byte(data), blackfriday.HtmlRenderer(htmlFlags, "", ""), mdExtensions)
	// Strip out bad HTML
	policy := bluemonday.UGCPolicy()
	policy.AllowAttrs("class", "id").Globally()
	outHTML := string(policy.SanitizeBytes(md))
	outHTML = markeddownReg.ReplaceAllString(outHTML, "$1")
	outHTML = strings.TrimRightFunc(outHTML, unicode.IsSpace)

	return outHTML
}

// ApplyBasicAccessibleMarkdown applies Markdown to the given data, rendering basic text formatting and preserving hard
// line breaks in HTML. It is meant for formatting text in small, multi-line UI elements, like user profile biographies.
func ApplyBasicAccessibleMarkdown(data []byte) string {
	mdExtensions := 0 |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_HARD_LINE_BREAK
	htmlFlags := 0 |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_SMARTYPANTS_DASHES

	// Generate Markdown
	md := blackfriday.Markdown([]byte(data), blackfriday.HtmlRenderer(htmlFlags, "", ""), mdExtensions)
	// Strip out bad HTML
	policy := bluemonday.UGCPolicy()
	policy.AllowAttrs("class", "id").Globally()
	policy.AllowAttrs("rel").OnElements("a")
	policy.RequireNoFollowOnLinks(false)
	outHTML := string(policy.SanitizeBytes(md))
	// Strip surrounding <p> tags that blackfriday adds
	outHTML = markeddownReg.ReplaceAllString(outHTML, "$1")
	outHTML = strings.TrimRightFunc(outHTML, unicode.IsSpace)

	return outHTML
}

// StripHTMLWithoutEscaping strips HTML tags with bluemonday's StrictPolicy, then unescapes the HTML
// entities added in by sanitizing the content.
func StripHTMLWithoutEscaping(content string) string {
	return html.UnescapeString(bluemonday.StrictPolicy().Sanitize(content))
}
