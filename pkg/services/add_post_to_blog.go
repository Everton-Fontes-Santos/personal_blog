package services

import (
	"errors"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/domain"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
)

// This Service execute the add of post in the blog
// in that case he will add in repo
type AddPostToBlog struct {
	PostRepository domain.PostRepository
}

func NewAddToBlogService(repo domain.PostRepository) ServiceInterface[any] {
	return &AddPostToBlog{
		PostRepository: repo,
	}
}

func (s *AddPostToBlog) Execute(opts ...any) (any, error) {
	var blog *entitys.Blog
	var post *entitys.Post

	for _, opt := range opts {
		switch v := opt.(type) {
		default:
			return nil, errors.New("the types must be Blog or Post")
		case *entitys.Post:
			post = v
		case *entitys.Blog:
			blog = v

		case entitys.Post:
			post = &v
		case entitys.Blog:
			blog = &v
		}
	}

	err := s.PostRepository.Save(*post)
	if err != nil {
		return nil, err
	}
	blog.AddPost(post)
	return nil, nil
}
