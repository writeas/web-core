// Package activitystreams provides all the basic ActivityStreams
// implementation needed for Write.as.
package activitystreams

import (
	"time"

	"github.com/writeas/web-core/id"
)

const (
	Namespace = "https://www.w3.org/ns/activitystreams"
	ToPublic  = "https://www.w3.org/ns/activitystreams#Public"
)

var (
	Extensions = map[string]string{}

	apMonetizationContext = "https://webmonetization.org/ns.jsonld"
)

// Activity describes actions that have either already occurred, are in the
// process of occurring, or may occur in the future.
type Activity struct {
	BaseObject
	Actor     string     `json:"actor"`
	Published time.Time  `json:"published,omitempty"`
	Updated   *time.Time `json:"updated,omitempty"`
	To        []string   `json:"to,omitempty"`
	CC        []string   `json:"cc,omitempty"`
	Object    *Object    `json:"object"`
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
	if o.Updated != nil && !o.Updated.IsZero() {
		a.Updated = o.Updated
	}
	return &a
}

// ActorActivity describes an action whose object is a full actor (Person),
// such as an Update to a profile's name, summary, or icon. It parallels
// Activity, which instead carries a content *Object.
type ActorActivity struct {
	BaseObject
	Actor   string     `json:"actor"`
	Updated *time.Time `json:"updated,omitempty"`
	To      []string   `json:"to,omitempty"`
	CC      []string   `json:"cc,omitempty"`
	Object  *Person    `json:"object"`
}

// NewUpdateActorActivity builds an Update activity whose object is the given
// actor (Person), used to federate profile changes such as the actor's name,
// summary, or icon. It addresses the public namespace and CCs the actor's
// followers collection. The activity gets a unique ID derived from the actor's
// IRI so remote instances don't dedupe repeated updates against one another.
func NewUpdateActorActivity(p *Person) *ActorActivity {
	a := ActorActivity{
		BaseObject: BaseObject{
			Context: []interface{}{
				Namespace,
			},
			ID:   p.ID + "#update-" + id.GenerateFriendlyRandomString(20),
			Type: "Update",
		},
		Actor:  p.ID,
		To:     []string{ToPublic},
		CC:     []string{p.Followers},
		Object: p,
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
	Published    time.Time         `json:"published,omitempty"`
	Updated      *time.Time        `json:"updated,omitempty"`
	Summary      *string           `json:"summary,omitempty"`
	InReplyTo    *string           `json:"inReplyTo,omitempty"`
	URL          string            `json:"url"`
	AttributedTo string            `json:"attributedTo,omitempty"`
	To           []string          `json:"to,omitempty"`
	CC           []string          `json:"cc,omitempty"`
	Name         string            `json:"name,omitempty"`
	Content      string            `json:"content,omitempty"`
	ContentMap   map[string]string `json:"contentMap,omitempty"`
	Tag          []Tag             `json:"tag,omitempty"`
	Attachment   []Attachment      `json:"attachment,omitempty"`
	Preview      *Object           `json:"preview,omitempty"`

	// Person
	Inbox             string     `json:"inbox,omitempty"`
	Outbox            string     `json:"outbox,omitempty"`
	Following         string     `json:"following,omitempty"`
	Followers         string     `json:"followers,omitempty"`
	PreferredUsername string     `json:"preferredUsername,omitempty"`
	Icon              *Image     `json:"icon,omitempty"`
	PublicKey         *PublicKey `json:"publicKey,omitempty"`
	Endpoints         *Endpoints `json:"endpoints,omitempty"`

	// Extensions
	Monetization string `json:"monetization,omitempty"`
}

// NewNoteObject creates a basic Note object that includes the public
// namespace in IRIs it's addressed to.
func NewNoteObject() *Object {
	o := Object{
		BaseObject: BaseObject{
			Type: "Note",
			Context: []interface{}{
				Namespace,
			},
		},
		To: []string{
			ToPublic,
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
			Context: []interface{}{
				Namespace,
			},
		},
		To: []string{
			ToPublic,
		},
	}
	return &o
}

// NewPersonObject creates a basic Person object.
func NewPersonObject() *Object {
	o := Object{
		BaseObject: BaseObject{
			Type: "Person",
			Context: []interface{}{
				Namespace,
			},
		},
	}
	return &o
}

func (o *Object) AddWebMonetization(wallet string) {
	o.Context = append(o.Context, apMonetizationContext)
	o.Monetization = wallet
}
