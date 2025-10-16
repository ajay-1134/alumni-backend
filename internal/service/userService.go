package service

import (
	"errors"

	"github.com/ajay-1134/alumni-backend/internal/domain"
	"github.com/ajay-1134/alumni-backend/internal/dto"
	"github.com/ajay-1134/alumni-backend/internal/ports/repository"
	"github.com/ajay-1134/alumni-backend/internal/ports/service"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) service.UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(req *dto.RegisterRequest) error {
	hashedPassword, _ := HashPassword(req.Password)
	u := registerRequestToUser(req)
	u.PasswordHash = hashedPassword
	return s.repo.Create(u)
}

func (s *userService) Login(email string, password string) (*domain.User, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil,err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return nil,errors.New("invalid credentials")
	}

	return u, nil
}

func (s *userService) UpdateDetails(userID uint, req *dto.UpdateUserRequest) error {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}
	updates := buildUpdates(req)

	err = s.repo.Update(user,updates)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteProfile(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *userService) GetAllUsers() ([]*dto.UserResponse, error) {
	users,err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var resp []*dto.UserResponse
	for _,user := range users {
		resp = append(resp, dto.ToUserResponse(user))
	}
	return resp,nil
}

func (s *userService) GetUser(id uint) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	resp := dto.ToUserResponse(user)
	return resp,nil
}

func (s *userService) LoginWithGoogle(googleUser *dto.GoogleUser) (*domain.User, error) {
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

func (s *userService) UserCount() (*int64, error) {
	return s.repo.UserCount()
}

func (s *userService) VerifiedUsersCount() (*int64,error) {
	return s.repo.VerifiedUsersCount()
}
