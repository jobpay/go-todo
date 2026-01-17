package valueobject

import (
	"errors"
	"regexp"
)

type Email string

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(email string) (Email, error) {
	if email == "" {
		return "", errors.New("email is required")
	}
	if !emailRegex.MatchString(email) {
		return "", errors.New("invalid email format")
	}
	if len(email) > 255 {
		return "", errors.New("email must be 255 characters or less")
	}
	return Email(email), nil
}

func (e Email) String() string {
	return string(e)
}

