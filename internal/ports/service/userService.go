package service

import (
	"github.com/ajay-1134/alumni-backend/internal/domain"
	"github.com/ajay-1134/alumni-backend/internal/dto"
)

type UserService interface {
	Register(*dto.RegisterRequest) error
	Login(email string, Password string) (*domain.User,error)
	LoginWithGoogle(googleUser *dto.GoogleUser) (*domain.User, error)
	GetAllUsers() ([]*dto.UserResponse, error)
	GetUser(id uint) (*dto.UserResponse, error)
	UpdateDetails(id uint, req *dto.UpdateUserRequest) error
	DeleteProfile(id uint) error
}
