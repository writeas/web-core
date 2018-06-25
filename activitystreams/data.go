package activitystreams

import "fmt"

type (
	BaseObject struct {
		Context []string `json:"@context,omitempty"`
		Type    string   `json:"type"`
		ID      string   `json:"id"`
	}

	PublicKey struct {
		ID           string `json:"id"`
		Owner        string `json:"owner"`
		PublicKeyPEM string `json:"publicKeyPem"`
		privateKey   []byte
	}

	Image struct {
		Type      string `json:"type"`
		MediaType string `json:"mediaType"`
		URL       string `json:"url"`
	}
)

type OrderedCollection struct {
	BaseObject
	TotalItems int    `json:"totalItems"`
	First      string `json:"first"`
	Last       string `json:"last,omitempty"`
}

func NewOrderedCollection(accountRoot string, items int) *OrderedCollection {
	oc := OrderedCollection{
		BaseObject: BaseObject{
			Context: []string{
				"https://www.w3.org/ns/activitystreams",
			},
			ID:   accountRoot + "/outbox",
			Type: "OrderedCollection",
		},
		First:      accountRoot + "/outbox?page=1",
		TotalItems: items,
	}
	return &oc
}

type OrderedCollectionPage struct {
	BaseObject
	TotalItems   int        `json:"totalItems"`
	PartOf       string     `json:"partOf"`
	Next         string     `json:"next,omitempty"`
	Prev         string     `json:"prev,omitempty"`
	OrderedItems []Activity `json:"orderedItems"`
}

func NewOrderedCollectionPage(accountRoot string, items, page int) *OrderedCollectionPage {
	ocp := OrderedCollectionPage{
		BaseObject: BaseObject{
			Context: []string{
				"https://www.w3.org/ns/activitystreams",
			},
			ID:   fmt.Sprintf("%s/outbox?page=%d", accountRoot, page),
			Type: "OrderedCollectionPage",
		},
		TotalItems: items,
		PartOf:     accountRoot + "/outbox",
		Next:       fmt.Sprintf("%s/outbox?page=%d", accountRoot, page+1),
	}
	return &ocp
}
