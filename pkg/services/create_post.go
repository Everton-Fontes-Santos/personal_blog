package services

import (
	"errors"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
)

type CreatePostService struct{}

func NewCreatePostService() ServiceInterface[entitys.Post] {
	return &CreatePostService{}
}

type PostInput struct {
	TITLE      string
	SLUG       string
	SMALL_DESC string
	TEXT       string
	AUTHOR     string
	CATEGORYS  []string
}

func defaultPostInput() PostInput {
	return PostInput{
		TITLE:      "Hello World",
		SLUG:       "hello-world",
		SMALL_DESC: "This is the first post of the blog c'mon",
		TEXT:       "This is the first post of the blog c'mon",
		CATEGORYS:  make([]string, 0),
	}
}

func (s *CreatePostService) Execute(opts ...any) (entitys.Post, error) {
	var inputOpts PostInput

	defaultOpt := defaultPostInput()

	if len(opts) > 0 {
		switch v := opts[0].(type) {
		default:
			return entitys.Post{}, errors.New("types must be PostInput")
		case PostInput:
			inputOpts = v
		case *PostInput:
			inputOpts = *v
		}

		if inputOpts.AUTHOR != "" {
			defaultOpt.AUTHOR = inputOpts.AUTHOR
		}
		if inputOpts.TEXT != "" {
			defaultOpt.TEXT = inputOpts.TEXT
		}
		if inputOpts.TITLE != "" {
			defaultOpt.TITLE = inputOpts.TITLE
		}
		if inputOpts.SLUG != "" {
			defaultOpt.SLUG = inputOpts.SLUG
		}
		if inputOpts.SMALL_DESC != "" {
			defaultOpt.SMALL_DESC = inputOpts.SMALL_DESC
		}
		// if len(inputOpts.CATEGORYS) != 0 {defaultOpt.CATEGORYS = inputOpts.CATEGORYS }
	}

	post, err := entitys.NewPost(
		entitys.WithOwner(defaultOpt.AUTHOR),
		entitys.WithText(defaultOpt.TEXT),
		entitys.WithTitle(defaultOpt.TITLE),
		entitys.WithSlug(defaultOpt.SLUG),
		entitys.WithSmallDesc(defaultOpt.SMALL_DESC),
	)
	if err != nil {
		return entitys.Post{}, err
	}
	return *post, nil
}
