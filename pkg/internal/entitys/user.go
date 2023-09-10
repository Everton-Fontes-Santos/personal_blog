package entitys

import "github.com/google/uuid"

// User is a entity thats represents a User for all Domains
// this entitys will get a id for unique comparations and
// a name and last name and age of the persor who he are
// and have a email as password
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
			ID:       id.String(),
			NAME:     name,
			LASTNAME: lastName,
			AGE:      29,
		},
		EMAIL:    email,
		PASSWORD: password,
	}, nil
}
