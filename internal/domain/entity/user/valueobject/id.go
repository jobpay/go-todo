package valueobject

import "errors"

type ID int

func NewID(id int) (ID, error) {
	if id <= 0 {
		return 0, errors.New("id must be greater than 0")
	}
	return ID(id), nil
}

func (i ID) Int() int {
	return int(i)
}

