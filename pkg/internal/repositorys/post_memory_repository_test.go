package repositorys_test

import (
	"testing"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/repositorys"
	"github.com/stretchr/testify/assert"
)

func TestPostMemoryRepository(t *testing.T) {
	repo := repositorys.NewPostMemoryRepository()
	post, _ := entitys.NewPost()
	_, err := repo.List()
	assert.Equal(t, error.Error(err), "don't have one post to return")
	err = repo.Save(*post)
	assert.NoError(t, err)
	posts, _ := repo.List()
	assert.NotEmpty(t, posts)
	assert.Equal(t, len(posts), 1)
	dbPost, err := repo.Get(post.ID)
	assert.NoError(t, err)
	assert.Equal(t, dbPost, *post)
}
