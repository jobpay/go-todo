package valueobject

import "errors"

type Name string

func NewName(name string) (Name, error) {
	if name == "" {
		return "", errors.New("name is required")
	}
	if len(name) > 100 {
		return "", errors.New("name must be 100 characters or less")
	}
	return Name(name), nil
}

func (n Name) String() string {
	return string(n)
}

