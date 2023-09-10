package services_test

import (
	"testing"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/repositorys"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBlogPost(t *testing.T) {
	repo := repositorys.NewPostMemoryRepository()
	post, _ := entitys.NewPost()
	blog := entitys.NewBlog("Efs Blog")
	add := services.NewAddToBlogService(repo)
	service := services.NewGetAllPostService(repo)
	add.Execute(post, blog)
	resp, err := service.Execute(blog)
	assert.NoError(t, err)
	assert.Equal(t, resp[0].ID, post.ID)

}
