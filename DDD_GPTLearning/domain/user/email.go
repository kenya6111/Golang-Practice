package user

import (
	"errors"
	"regexp"
)

type Email string

func NewEmail(value string) (Email, error) {
	if !regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`).MatchString(value) {
		return "", errors.New("invalid email format")
	}
	return Email(value), nil
}

func (e Email) String() string {
	return string(e)
}
