package services

import (
	"errors"
	"time"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/domain"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
)

// This Service execute the add of post in the blog
// in that case he will add in repo
type GetAllPostsByDateService struct {
	PostRepository domain.PostRepository
}

func NewGetAllPostsByDateService(repo domain.PostRepository) ServiceInterface[[]entitys.Post] {
	return &GetAllPostsByDateService{
		PostRepository: repo,
	}
}

func (s *GetAllPostsByDateService) Execute(opts ...any) ([]entitys.Post, error) {

	var blog entitys.Blog
	var posts []entitys.Post
	var date time.Time

	for _, opt := range opts {
		switch v := opt.(type) {
		default:
			return []entitys.Post{}, errors.New("types must be Blog")
		case time.Time:
			date = v
		case string:
			dat, err := time.Parse(time.RFC3339, v)
			if err != nil {
				return []entitys.Post{}, err
			}
			date = dat
		case entitys.Blog:
			blog = v
		case *entitys.Blog:
			blog = *v
		}
	}

	posts, err := s.PostRepository.List()
	if err != nil {
		return []entitys.Post{}, err
	}
	blog.UpdatePosts(posts)

	posts, err = blog.GetPostByDate(date)
	if err != nil {
		return []entitys.Post{}, err
	}

	return posts, nil
}
