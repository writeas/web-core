package errors

import (
	"github.com/writeas/impart"
	"net/http"
)

var (
	// Commonly returned HTTP errors
	ErrBadFormData      = impart.HTTPError{http.StatusBadRequest, "Expected valid form data."}
	ErrBadJSON          = impart.HTTPError{http.StatusBadRequest, "Expected valid JSON object."}
	ErrBadJSONArray     = impart.HTTPError{http.StatusBadRequest, "Expected valid JSON array."}
	ErrBadAccessToken   = impart.HTTPError{http.StatusUnauthorized, "Invalid access token."}
	ErrNoAccessToken    = impart.HTTPError{http.StatusBadRequest, "Authorization token required."}
	ErrNotLoggedIn      = impart.HTTPError{http.StatusUnauthorized, "Not logged in."}
	ErrBadRequestedType = impart.HTTPError{http.StatusNotAcceptable, "Bad requested Content-Type."}

	// Post operation errors
	ErrPostNoUpdatableVals  = impart.HTTPError{http.StatusBadRequest, "Supply some properties to update."}
	ErrNoPublishableContent = impart.HTTPError{http.StatusBadRequest, "Supply something to publish."}

	// Internal errors
	ErrInternalGeneral       = impart.HTTPError{http.StatusInternalServerError, "The humans messed something up. They've been notified."}
	ErrInternalCookieSession = impart.HTTPError{http.StatusInternalServerError, "Could not get cookie session."}

	// User errors
	ErrUserNotFound      = impart.HTTPError{http.StatusNotFound, "User doesn't exist."}
	ErrUserNotFoundEmail = impart.HTTPError{http.StatusNotFound, "Please enter your username instead of your email address."}
	ErrUsernameTaken     = impart.HTTPError{http.StatusConflict, "Username is already taken."}
)
