package activitystreams

type (
	BaseObject struct {
		Context []string `json:"@context"`
		Type    string   `json:"type"`
		ID      string   `json:"id"`
	}

	PublicKey struct {
		ID           string `json:"id"`
		Owner        string `json:"owner"`
		PublicKeyPEM string `json:"publicKeyPem"`
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
