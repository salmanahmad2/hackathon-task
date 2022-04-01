package server

import (
	"github.com/google/uuid"
)

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func TokenValidator(token string) bool {
	var (
		isEmpty        = false
		hasTokenLength = false
	)
	if len(token) > 0 {
		isEmpty = true
	}
	if len(token) == 6 {
		hasTokenLength = true
	}
	return isEmpty && hasTokenLength
}
