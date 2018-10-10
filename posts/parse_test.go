package posts

import (
	"testing"
)

type titleTest struct {
	in, title, body string
}

func TestExtractTitle(t *testing.T) {
	tests := []titleTest{
		{`# Hello World
This is my post`, "Hello World", "This is my post"},
		{"No title", "", "No title"},
		{`Not explicit title

It's not explicit.
Yep.`, "", `Not explicit title

It's not explicit.
Yep.`},
		{"# Only a title", "", "# Only a title"},
	}

	for _, test := range tests {
		title, body := ExtractTitle(test.in)
		if title != test.title {
			t.Fatalf("Wanted title '%s', got '%s'", test.title, title)
		}
		if body != test.body {
			t.Fatalf("Wanted body '%s', got '%s'", test.body, body)
		}
	}
}
