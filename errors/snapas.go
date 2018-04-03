package errors

import (
	"github.com/writeas/impart"
	"net/http"
)

var (
	ErrInvalidFeature = impart.HTTPError{http.StatusBadRequest, "Not a valid feature."}
)
