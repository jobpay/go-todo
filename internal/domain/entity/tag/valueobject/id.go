package valueobject

import "errors"

type ID int

func NewID(id int) (ID, error) {
	if id < 0 {
		return 0, errors.New("id must be positive")
	}
	return ID(id), nil
}

func (id ID) Int() int {
	return int(id)
}

