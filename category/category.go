// Package category supports post categories
package category

import (
	"errors"
	"github.com/writeas/slug"
)

var (
	ErrNotFound = errors.New("category doesn't exist")
)

// Category represents a post tag with additional metadata, like a title and slug.
type Category struct {
	ID      int64  `json:"-"`
	Hashtag string `json:"hashtag"`
	Slug    string `json:"slug"`
	Title   string `json:"title"`
}

// NewCategory creates a Category you can insert into the database, based on a hashtag. It automatically breaks up the
// hashtag by words, based on capitalization, for both the title and a URL-friendly slug.
func NewCategory(hashtag string) *Category {
	title := titleFromHashtag(hashtag)
	return &Category{
		Hashtag: hashtag,
		Slug:    slug.Make(title),
		Title:   title,
	}
}
