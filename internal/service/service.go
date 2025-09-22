package service

import (
	"errors"
	"fmt"

	"github.com/ajay-1134/alumni-backend/internal/domain"
	"github.com/ajay-1134/alumni-backend/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(req *dto.RegisterRequest) error {
	hashedPassword, _ := HashPassword(req.Password)
	u := registerRequestToUser(req)
	u.PasswordHash = hashedPassword
	return s.repo.Create(u)
}

func (s *UserService) Login(email string, password string) (*domain.User, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	fmt.Printf("password: %v\n", password)
	fmt.Printf("u.PasswordHash: %v\n", u.PasswordHash)
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return u, nil
}

func (s *UserService) UpdateDetails(userID uint, req *dto.UpdateUserRequest) (*domain.User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.Update(user, req)
}

func (s *UserService) DeleteProfile(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetUser(id uint) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) LoginWithGoogle(googleUser *dto.GoogleUser) (*domain.User, error) {
	email := googleUser.Email
	user := googleUserToUser(googleUser)
	_, err := s.repo.FindByEmail(email)
	if err != nil {
		err = s.repo.Create(user)
		if err != nil {
			return nil, err
		}
	}

	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return u, nil
}
