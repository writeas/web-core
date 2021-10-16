package category

import "testing"

func TestTitleFromHashtag(t *testing.T) {
	tests := []struct {
		name     string
		hashtag  string
		expTitle string
	}{
		{"proper noun", "Jane", "jane"},
		{"full name", "JaneDoe", "jane doe"},
		{"us words", "unitedStates", "united states"},
		{"usa", "USA", "usa"},
		{"us monoword", "unitedstates", "unitedstates"},
		{"100dto", "100DaysToOffload", "100 days to offload"},
		{"iphone", "iPhone", "iphone"},
		{"ilike", "iLikeThis", "i like this"},
		{"abird", "aBird", "a bird"},
		{"all caps", "URGENT", "urgent"},
		{"smartphone", "スマートフォン", "スマートフォン"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := titleFromHashtag(test.hashtag)
			if res != test.expTitle {
				t.Fatalf("#%s: got '%s' expected '%s'", test.hashtag, res, test.expTitle)
			}
		})
	}
}

func TestHashtagFromTitle(t *testing.T) {
	tests := []struct {
		name       string
		title      string
		expHashtag string
	}{
		{"proper noun", "Jane", "Jane"},
		{"full name", "Jane Doe", "JaneDoe"},
		{"us upper words", "United States", "UnitedStates"},
		{"us lower words", "united states", "unitedStates"},
		{"usa", "USA", "USA"},
		{"100dto", "100 Days To Offload", "100DaysToOffload"},
		{"iphone", "iPhone", "iPhone"},
		{"ilike", "I like this", "ILikeThis"},
		{"abird", "a Bird", "aBird"},
		{"all caps", "URGENT", "URGENT"},
		{"punctuation", "John’s Stories", "JohnsStories"},
		{"smartphone", "スマートフォン", "スマートフォン"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := HashtagFromTitle(test.title)
			if res != test.expHashtag {
				t.Fatalf("%s: got '%s' expected '%s'", test.title, res, test.expHashtag)
			}
		})
	}
}
