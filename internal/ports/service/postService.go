package service

import (
	"github.com/ajay-1134/alumni-backend/internal/dto"
)

type PostService interface {
	CreatePost(*dto.PostDto) error
	GetAllPosts() ([]*dto.PostDto,error)
	GetAllPostsWithUserId(userId uint) ([]*dto.PostDto,error)
	UpdatePost(id uint, postDto *dto.PostDto) error
	DeletePost(id uint) error
}
