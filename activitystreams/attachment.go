package activitystreams

type Attachment struct {
	Type      AttachmentType `json:"type"`
	URL       string         `json:"url"`
	MediaType string         `json:"mediaType"`
	Name      string         `json:"name"`
}

type AttachmentType string

const AttachImage AttachmentType = "Image"
