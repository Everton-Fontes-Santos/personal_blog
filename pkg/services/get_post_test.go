package services_test

import (
	"testing"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/repositorys"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestGetBlogPost(t *testing.T) {
	repo := repositorys.NewPostMemoryRepository()
	post, _ := entitys.NewPost()
	blog := entitys.NewBlog("Efs Blog")
	add := services.NewAddToBlogService(repo)
	service := services.NewGetByIDService(repo)
	add.Execute(post, blog)

	resp, err := service.Execute(blog, post.ID)
	assert.NoError(t, err)
	assert.Equal(t, resp.ID, post.ID)

}
