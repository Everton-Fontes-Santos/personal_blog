package entitys

import (
	"time"

	"github.com/google/uuid"
)

// POst is a entity thats represents a Post for all Domains
// this entity will get a id for unique comparations and
// a title and slug for url, small desck to show on page and the text
type Post struct {
	ID         uuid.UUID `json:"id"`
	TITLE      string    `json:"title"`
	SLUG       string    `json:"slug"`
	SMALL_DESC string    `json:"small_desc"`
	TEXT       string    `json:"text"`
	CREATED_AT time.Time `json:"created_at"`
	UPDATED_AT time.Time `json:"updated_at"`
	OWNER_ID   uuid.UUID `json:"owner_id"`
}
