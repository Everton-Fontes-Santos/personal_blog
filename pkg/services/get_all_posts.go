package services

import (
	"errors"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/domain"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
)

// This Service execute the add of post in the blog
// in that case he will add in repo
type GetAllPostsService struct {
	PostRepository domain.PostRepository
}

func NewGetAllPostService(repo domain.PostRepository) ServiceInterface[[]entitys.Post] {
	return &GetAllPostsService{
		PostRepository: repo,
	}
}

func (s *GetAllPostsService) Execute(opts ...any) ([]entitys.Post, error) {

	var blog entitys.Blog
	var posts []entitys.Post

	for _, opt := range opts {
		switch v := opt.(type) {
		default:
			return []entitys.Post{}, errors.New("types must be Blog")
		case entitys.Blog:
			blog = v
		case *entitys.Blog:
			blog = *v
		}
	}

	posts, err := blog.GetAllPosts()
	if err != nil {
		posts, err = s.PostRepository.List()
		if err != nil {
			return []entitys.Post{}, err
		}
	}
	return posts, nil
}
