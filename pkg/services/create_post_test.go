package services_test

import (
	"testing"

	"github.com/Everton-Fontes-Santos/personal_blog/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	service := services.NewCreatePostService()

	text := "This will be the text"
	props := services.PostInput{
		TEXT: text,
	}
	resp, err := service.Execute(props)
	assert.NoError(t, err)
	assert.NotEqual(t, resp, nil)
	assert.Equal(t, text, resp.TEXT)
	assert.Equal(t, "Hello World", resp.TITLE)
}
