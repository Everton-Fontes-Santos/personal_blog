package domain

import (
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
)

type PostRepository interface {
	Save(post entitys.Post) error
	Delete(id string) error
	Get(id string) (entitys.Post, error)
	List() ([]entitys.Post, error)
}
