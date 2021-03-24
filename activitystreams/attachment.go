package activitystreams

import (
	"mime"
	"strings"
)

type Attachment struct {
	Type      AttachmentType `json:"type"`
	URL       string         `json:"url"`
	MediaType string         `json:"mediaType"`
	Name      string         `json:"name"`
}

type AttachmentType string

const AttachImage AttachmentType = "Image"

// NewImageAttachment creates a new Attachment from the given URL, setting the
// correct type and automatically detecting the MediaType based on the file
// extension.
func NewImageAttachment(url string) *Attachment {
	var imgType string
	extIdx := strings.LastIndexByte(url, '.')
	if extIdx > -1 {
		imgType = mime.TypeByExtension(url[extIdx:])
	}
	return &Attachment{
		Type:      AttachImage,
		URL:       url,
		MediaType: imgType,
	}
}
