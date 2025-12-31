package valueobject

import "errors"

type Title string

func NewTitle(title string) (Title, error) {
	if title == "" {
		return "", errors.New("title is required")
	}
	if len(title) > 100 {
		return "", errors.New("title must be 100 characters or less")
	}
	return Title(title), nil
}

func (t Title) String() string {
	return string(t)
}

