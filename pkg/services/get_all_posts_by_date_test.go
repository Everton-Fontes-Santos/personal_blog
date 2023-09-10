package services_test

import (
	"testing"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/entitys"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/internal/repositorys"
	"github.com/Everton-Fontes-Santos/personal_blog/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBlogPostByDate(t *testing.T) {
	repo := repositorys.NewPostMemoryRepository()

	dateString := "2023-08-08T11:00:00Z"
	post, _ := entitys.NewPost(entitys.WithCreatedAt(dateString))

	dateString2 := "2023-08-08T12:00:00Z"
	post2, _ := entitys.NewPost(entitys.WithCreatedAt(dateString2))

	blog := entitys.NewBlog("Efs Blog")
	add := services.NewAddToBlogService(repo)
	service := services.NewGetAllPostsByDateService(repo)
	add.Execute(post, blog)
	add.Execute(post2, blog)

	resp, err := service.Execute(blog, dateString)
	assert.NoError(t, err)
	assert.Equal(t, resp[0].ID, post.ID)
	assert.Equal(t, resp[1].ID, post2.ID)

}
