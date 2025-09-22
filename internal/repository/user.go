package repository

import (
	"github.com/ajay-1134/alumni-backend/internal/domain"
	"github.com/ajay-1134/alumni-backend/internal/dto"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *domain.User, req *dto.UpdateUserRequest) (*domain.User, error) {

	if err := r.db.Model(&user).Updates(req).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}

func (r *userRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
