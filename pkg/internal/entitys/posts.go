package entitys

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type PostOps struct {
	ID         string      `json:"id"`
	TITLE      string      `json:"title"`
	SLUG       string      `json:"slug"`
	SMALL_DESC string      `json:"small_desc"`
	TEXT       string      `json:"text"`
	CREATED_AT time.Time   `json:"created_at"`
	UPDATED_AT time.Time   `json:"updated_at"`
	AUTHOR     string      `json:"author"`
	CATEGORYS  []*Category `json:"categorys"`
}

// POst is a entity thats represents a Post for all Domains
// this entity will get a id for unique comparations and
// a title and slug for url, small desck to show on page and the text
type Post struct {
	PostOps
}

type PostOptsFunc func(*PostOps)

func defaultPostOpts() (PostOps, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return PostOps{}, errors.New("errors on creating a uuid for post")
	}
	return PostOps{
		ID:         id.String(),
		CREATED_AT: time.Now(),
		CATEGORYS:  make([]*Category, 0),
	}, nil
}

func WithCategory(cat *Category) PostOptsFunc {
	return func(opts *PostOps) {
		opts.CATEGORYS = append(opts.CATEGORYS, cat)
	}
}

func WithOwner(id string) PostOptsFunc {
	return func(opts *PostOps) {
		opts.AUTHOR = id
	}
}

func WithTitle(title string) PostOptsFunc {
	return func(opts *PostOps) {
		opts.TITLE = title
	}
}

func WithText(text string) PostOptsFunc {
	return func(opts *PostOps) {
		opts.TEXT = text
	}
}

func WithSlug(slug string) PostOptsFunc {
	return func(opts *PostOps) {
		opts.SLUG = slug
	}
}

func WithSmallDesc(desc string) PostOptsFunc {
	return func(opts *PostOps) {
		opts.SMALL_DESC = desc
	}
}

func NewPost(opts ...PostOptsFunc) (*Post, error) {
	defaultOpts, err := defaultPostOpts()
	if err != nil {
		return &Post{}, err
	}

	for _, fn := range opts {
		fn(&defaultOpts)
	}

	return &Post{
		PostOps: defaultOpts,
	}, nil
}

func (p *Post) SetText(text string) {
	p.TEXT = text
	p.UPDATED_AT = time.Now()
}

func (p *Post) SetSmallDesc(text string) {
	p.SMALL_DESC = text
	p.UPDATED_AT = time.Now()
}

func (p *Post) SetSlug(text string) {
	p.SLUG = text
	p.UPDATED_AT = time.Now()
}

func (p *Post) SetTitle(text string) {
	p.TITLE = text
	p.UPDATED_AT = time.Now()
}
