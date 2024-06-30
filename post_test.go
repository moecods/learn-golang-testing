package tutorialgolangtesting

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostSerive_GetPostTitle(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := &PostService{repo: mockRepo}

	mockRepo.On("GetPost", 1).Return(&Post{ID: 1, Title: "Me Before You"}, nil)
	title, err := service.GetPostTitle(1)
	assert.NoError(t, err)
	assert.Equal(t, "Me Before You", title)

	mockRepo.On("GetPost", 2).Return(nil, errors.New("post not found"))
	_, err = service.GetPostTitle(2)
    assert.Error(t, err)
}