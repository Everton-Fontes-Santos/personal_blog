package entitys

import (
	"errors"
	"time"
)

type Blog struct {
	NAME   string
	posts  []*Post
	admins []*User
}

func NewBlog(name string) *Blog {
	return &Blog{
		NAME:   name,
		posts:  make([]*Post, 0),
		admins: make([]*User, 0),
	}
}

func (b *Blog) AddAdmin(user *User) error {
	for _, _user := range b.admins {
		if user.EMAIL == _user.EMAIL {
			return errors.New("this user already is a admin")
		}
	}
	b.admins = append(b.admins, user)
	return nil
}

func (b *Blog) GetAdmins() ([]*User, error) {
	if len(b.admins) <= 0 {
		return []*User{}, errors.New("don't have any user to get")
	}
	return b.admins, nil
}

func (b *Blog) GetAllPosts() ([]*Post, error) {
	if len(b.posts) <= 0 {
		return []*Post{}, errors.New("don't have any post to get")
	}
	return b.posts, nil
}

func (b *Blog) GetPostByDate(date time.Time) ([]*Post, error) {
	posts, err := b.GetAllPosts()
	if err != nil {
		return posts, err
	}
	_posts := make([]*Post, 0)
	for _, post := range posts {
		if date.Equal(post.CREATED_AT) {
			_posts = append(_posts, post)
		}
	}
	if len(_posts) == 0 {
		return []*Post{}, errors.New("don't have any post to get")
	}
	return _posts, nil
}

func (b *Blog) AddPost(post *Post) {
	b.posts = append(b.posts, post)
}
