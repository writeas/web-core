package activitystreams

import (
	"time"
)

const (
	toPublic = "https://www.w3.org/ns/activitystreams#Public"
)

type Activity struct {
	BaseObject
	Actor     string    `json:"actor"`
	Published time.Time `json:"published"`
	To        []string  `json:"to"`
	CC        []string  `json:"cc"`
	Object    *Object   `json:"object"`
}

func NewCreateActivity(o *Object) *Activity {
	a := Activity{
		BaseObject: BaseObject{
			ID:   o.ID + "/activity",
			Type: "Create",
		},
		Actor:     o.AttributedTo,
		Published: o.Published,
		To:        o.To,
		CC:        o.CC,
		Object:    o,
	}
	return &a
}

type Object struct {
	BaseObject
	Published    time.Time         `json:"published"`
	Summary      *string           `json:"summary"`
	InReplyTo    *string           `json:"inReplyTo"`
	URL          string            `json:"url"`
	AttributedTo string            `json:"attributedTo"`
	To           []string          `json:"to"`
	CC           []string          `json:"cc"`
	Content      string            `json:"content"`
	ContentMap   map[string]string `json:"contentMap"`
}

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
