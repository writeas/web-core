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

// NewCategoryFromPartial creates a Category from a partially-populated Category, such as when a user initially creates
// one.
func NewCategoryFromPartial(cat *Category) *Category {
	newCat := &Category{
		Hashtag: cat.Hashtag,
	}
	// Create title from hashtag, if none supplied
	if cat.Title == "" {
		newCat.Title = titleFromHashtag(cat.Hashtag)
	} else {
		newCat.Title = cat.Title
	}
	// Create slug from title, if none supplied; otherwise ensure slug is valid
	if cat.Slug == "" {
		newCat.Slug = slug.Make(newCat.Title)
	} else {
		newCat.Slug = slug.Make(cat.Slug)
	}
	return newCat
}
