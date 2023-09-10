package entitys

import (
	"errors"

	"github.com/google/uuid"
)

type Category struct {
	ID   string
	NAME string
}

func NewCategory(name string) (*Category, error) {

	id, err := uuid.NewUUID()
	if err != nil {
		return &Category{}, errors.New("errors on creating a uuid for category")
	}

	return &Category{
		ID:   id.String(),
		NAME: name,
	}, nil
}
