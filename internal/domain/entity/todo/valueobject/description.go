package valueobject

import "errors"

type Description string

func NewDescription(desc string) (Description, error) {
	if len(desc) > 500 {
		return "", errors.New("description must be 500 characters or less")
	}
	return Description(desc), nil
}

func (d Description) String() string {
	return string(d)
}
