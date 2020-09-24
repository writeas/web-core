package auth

import (
	uuid "github.com/gofrs/uuid"
	"github.com/writeas/web-core/log"
	"strings"
)

// GetToken parses out the user token from either an Authorization header or simply passed in.
func GetToken(header string) []byte {
	var accessToken []byte
	token := header
	if len(header) > 0 {
		f := strings.Fields(header)
		if len(f) == 2 && f[0] == "Token" {
			token = f[1]
		}
	}
	t, err := uuid.FromString(token)
	if err != nil {
		log.Error("Couldn't parseHex on '%s': %v", accessToken, err)
	} else {
		accessToken = t[:]
	}
	return accessToken
}

// GetHeaderToken parses out the user token from an Authorization header.
func GetHeaderToken(header string) []byte {
	var accessToken []byte
	if len(header) > 0 {
		f := strings.Fields(header)
		if len(f) == 2 && f[0] == "Token" {
			t, err := uuid.FromString(f[1])
			if err != nil {
				log.Error("Couldn't parseHex on '%s': %v", accessToken, err)
			} else {
				accessToken = t[:]
			}
		}
	}
	return accessToken
}
