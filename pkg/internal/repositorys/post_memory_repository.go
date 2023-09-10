package repositorys

import (
	"errors"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
)

type PostMemoryDB struct {
	Posts []*entitys.Post
}

type PostMemoryRepository struct {
	db PostMemoryDB
}

func NewPostMemoryRepository() *PostMemoryRepository {
	return &PostMemoryRepository{
		db: PostMemoryDB{
			Posts: make([]*entitys.Post, 0),
		},
	}
}

func (p *PostMemoryRepository) Get(id string) (*entitys.Post, error) {
	for _, post := range p.db.Posts {
		if post.ID == id {
			return post, nil
		}
	}
	return &entitys.Post{}, errors.New("don't have this post on db")
}

func (p *PostMemoryRepository) Delete(id string) error {
	for idx, post := range p.db.Posts {
		if post.ID == id {
			p.db.Posts = append(p.db.Posts[:idx], p.db.Posts[idx+1:]...)
			return nil
		}
	}
	return errors.New("don't have this post to delete")
}

func (p *PostMemoryRepository) List() ([]*entitys.Post, error) {
	if len(p.db.Posts) == 0 {
		return []*entitys.Post{}, errors.New("don't have one post to return")
	}
	return p.db.Posts, nil
}

func (p *PostMemoryRepository) Save(post *entitys.Post) error {
	dbPost, err := p.Get(post.ID)
	if err != nil {
		p.db.Posts = append(p.db.Posts, post)
		return nil
	}
	dbPost.AUTHOR = post.AUTHOR
	dbPost.CATEGORYS = post.CATEGORYS
	dbPost.CREATED_AT = post.CREATED_AT
	dbPost.SetSlug(post.SLUG)
	dbPost.SetSmallDesc(post.SMALL_DESC)
	dbPost.SetTitle(post.TITLE)
	dbPost.SetText(post.TEXT)
	return nil
}
