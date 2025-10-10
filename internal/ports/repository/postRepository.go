package repository

import "github.com/ajay-1134/alumni-backend/internal/domain"


type PostRepository interface {
	Create(post *domain.Post) error
	GetAll() ([]*domain.Post,error)
	GetAllWithUserId(userId uint) ([]*domain.Post,error)
	GetPostById(id uint) (*domain.Post,error)
	Update(post *domain.Post, updates *map[string]interface{}) error
	Delete(postId uint) error
}