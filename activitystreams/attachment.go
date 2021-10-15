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

const (
	AttachImage    AttachmentType = "Image"
	AttachDocument AttachmentType = "Document"
)

// NewImageAttachment creates a new Attachment from the given URL, setting the
// correct type and automatically detecting the MediaType based on the file
// extension.
func NewImageAttachment(url string) Attachment {
	return newAttachment(url, AttachImage)
}

// NewDocumentAttachment creates a new Attachment from the given URL, setting the
// correct type and automatically detecting the MediaType based on the file
// extension.
func NewDocumentAttachment(url string) Attachment {
	return newAttachment(url, AttachDocument)
}

func newAttachment(url string, attachType AttachmentType) Attachment {
	var fileType string
	extIdx := strings.LastIndexByte(url, '.')
	if extIdx > -1 {
		fileType = mime.TypeByExtension(url[extIdx:])
	}
	return Attachment{
		Type:      attachType,
		URL:       url,
		MediaType: fileType,
	}
}
