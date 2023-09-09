package entitys

import "github.com/google/uuid"

type User struct {
	*Person
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}

func NewUser(
	name string,
	lastName string,
	email string,
	password string,
) (*User, error) {

	id, err := uuid.NewUUID()
	if err != nil {
		return &User{}, err
	}

	return &User{
		Person: &Person{
			ID:       id,
			NAME:     name,
			LASTNAME: lastName,
			AGE:      29,
		},
		EMAIL:    email,
		PASSWORD: password,
	}, nil
}
