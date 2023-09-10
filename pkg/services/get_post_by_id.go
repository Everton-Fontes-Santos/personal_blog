package services

import (
	"errors"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/domain"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
)

// This Service execute the add of post in the blog
// in that case he will add in repo
type GetPostByIDService struct {
	PostRepository domain.PostRepository
}

func (s *GetPostByIDService) Execute(opts ...any) (any, error) {

	var blog entitys.Blog
	var id string

	for _, opt := range opts {
		switch v := opt.(type) {
		default:
			return nil, errors.New("types must be Blog or Post")
		case string:
			id = v
		case entitys.Blog:
			blog = v
		case *entitys.Blog:
			blog = *v
		}
	}

	posts, err := blog.GetAllPosts()
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		if post.ID == id {
			return post, nil
		}
	}
	post, err := s.PostRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}
