package services_test

import (
	"testing"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/repositorys"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestAddPostTOBlog(t *testing.T) {
	repo := repositorys.NewPostMemoryRepository()
	post, _ := entitys.NewPost()
	blog := entitys.NewBlog("Efs Blog")
	service := services.NewAddToBlogService(repo)
	resp, err := service.Execute(post, blog)
	assert.NoError(t, err)
	assert.Equal(t, resp, nil)
	// assert.Equal(t, error.Error(err), "don't have one post to return")

}
