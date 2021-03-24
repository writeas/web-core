// Package activitystreams provides all the basic ActivityStreams
// implementation needed for Write.as.
package activitystreams

import (
	"time"
)

const (
	Namespace = "https://www.w3.org/ns/activitystreams"
	toPublic  = "https://www.w3.org/ns/activitystreams#Public"
)

var Extensions = map[string]string{}

// Activity describes actions that have either already occurred, are in the
// process of occurring, or may occur in the future.
type Activity struct {
	BaseObject
	Actor     string    `json:"actor"`
	Published time.Time `json:"published,omitempty"`
	To        []string  `json:"to,omitempty"`
	CC        []string  `json:"cc,omitempty"`
	Object    *Object   `json:"object"`
}

type FollowActivity struct {
	BaseObject
	Actor     string    `json:"actor"`
	Published time.Time `json:"published,omitempty"`
	To        []string  `json:"to,omitempty"`
	CC        []string  `json:"cc,omitempty"`
	Object    string    `json:"object"`
}

// NewCreateActivity builds a basic Create activity that includes the given
// Object and the Object's AttributedTo property as the Actor.
func NewCreateActivity(o *Object) *Activity {
	a := Activity{
		BaseObject: BaseObject{
			Context: []interface{}{
				Namespace,
				Extensions,
			},
			ID:   o.ID,
			Type: "Create",
		},
		Actor:     o.AttributedTo,
		Object:    o,
		Published: o.Published,
	}
	return &a
}

// NewUpdateActivity builds a basic Update activity that includes the given
// Object and the Object's AttributedTo property as the Actor.
func NewUpdateActivity(o *Object) *Activity {
	a := Activity{
		BaseObject: BaseObject{
			Context: []interface{}{
				Namespace,
				Extensions,
			},
			ID:   o.ID,
			Type: "Update",
		},
		Actor:     o.AttributedTo,
		Object:    o,
		Published: o.Published,
	}
	return &a
}

// NewDeleteActivity builds a basic Delete activity that includes the given
// Object and the Object's AttributedTo property as the Actor.
func NewDeleteActivity(o *Object) *Activity {
	a := Activity{
		BaseObject: BaseObject{
			Context: []interface{}{
				Namespace,
			},
			ID:   o.ID,
			Type: "Delete",
		},
		Actor:  o.AttributedTo,
		Object: o,
	}
	return &a
}

// NewFollowActivity builds a basic Follow activity.
func NewFollowActivity(actorIRI, followeeIRI string) *FollowActivity {
	a := FollowActivity{
		BaseObject: BaseObject{
			Context: []interface{}{
				Namespace,
			},
			Type: "Follow",
		},
		Actor:  actorIRI,
		Object: followeeIRI,
	}
	return &a
}

// Object is the primary base type for the Activity Streams vocabulary.
type Object struct {
	BaseObject
	Published    time.Time         `json:"published"`
	Summary      *string           `json:"summary,omitempty"`
	InReplyTo    *string           `json:"inReplyTo"`
	URL          string            `json:"url"`
	AttributedTo string            `json:"attributedTo"`
	To           []string          `json:"to"`
	CC           []string          `json:"cc,omitempty"`
	Name         string            `json:"name,omitempty"`
	Content      string            `json:"content"`
	ContentMap   map[string]string `json:"contentMap,omitempty"`
	Tag          []Tag             `json:"tag"`
	Attachment   []Attachment      `json:"attachment,omitempty"`

	// Extensions
	// NOTE: add extensions here
}

// NewNoteObject creates a basic Note object that includes the public
// namespace in IRIs it's addressed to.
func NewNoteObject() *Object {
	o := Object{
		BaseObject: BaseObject{
			Type: "Note",
		},
		To: []string{
			toPublic,
		},
	}
	return &o
}

// NewArticleObject creates a basic Article object that includes the public
// namespace in IRIs it's addressed to.
func NewArticleObject() *Object {
	o := Object{
		BaseObject: BaseObject{
			Type: "Article",
		},
		To: []string{
			toPublic,
		},
	}
	return &o
}
