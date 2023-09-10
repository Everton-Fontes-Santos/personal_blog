package entitys

// Person is a entity thats represents a Person for all Domains
// this entitys will get a id for unique comparations and
// a name and last name and age
type Person struct {
	ID       string `json:"id"`
	NAME     string `json:"name"`
	LASTNAME string `json:"last_name"`
	AGE      int16  `json:"age"`
}
