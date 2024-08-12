package user

import (
	"errors"
	"strings"
)

type Username string

var (
	ErrEmptyUsername = errors.New("empty username supplied")
)

func NewUsername(value string) (Username, error) {
	un := strings.TrimSpace(value)
	// validate if un is empty
	if un == "" {
		return Username(""), ErrEmptyUsername
	}
	return Username(value), nil
}
