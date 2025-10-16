package repository

import "github.com/ajay-1134/alumni-backend/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByID(id uint) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	Update(user *domain.User, updates *map[string]interface{}) error
	Delete(id uint) error
	UserCount() (*int64,error)
	VerifiedUsersCount() (*int64,error)
}
