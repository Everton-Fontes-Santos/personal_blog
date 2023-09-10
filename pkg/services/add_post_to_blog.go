package services

import (
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/domain"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
)

// This Service execute the add of post in the blog
// in that case he will add in repo
type AddPostToBlog struct {
	PostRepository domain.PostRepository
}

func (s *AddPostToBlog) Execute(blog *entitys.Blog, post entitys.Post) (any, error) {
	err := s.PostRepository.Save(post)
	if err != nil {
		return nil, err
	}
	blog.AddPost(&post)
	return nil, nil
}
