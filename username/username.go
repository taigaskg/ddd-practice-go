package username

import (
	"errors"
	"unicode/utf8"
)

type UserName struct {
	value string
}

func New(value string) (*UserName, error) {
	if value == "" {
		return nil, errors.New("value must not be empty")
	}
	if utf8.RuneCountInString(value) < 3 {
		return nil, errors.New("value must be more than 3 characters.")
	}
	return &UserName{value}, nil
}
