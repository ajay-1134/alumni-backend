package repository

import (
	"log"

	"github.com/ajay-1134/alumni-backend/internal/domain"
	"github.com/ajay-1134/alumni-backend/internal/ports/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *domain.User, updates *map[string]interface{}) error {
	if err := r.db.Model(&user).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user *domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("error occured while finding user with email %s",email)
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user *domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		log.Printf("error occured while finding user with id %v",id)
		return nil,err
	}
	return user, nil
}

func (r *userRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}

func (r *userRepository) UserCount() (*int64, error) {
	var totalUsers int64

	err := r.db.Model(&domain.User{}).Count(&totalUsers).Error
	if err != nil {
		log.Printf("failed to count total users")
		return nil,err
	}
	return &totalUsers,nil
}

func (r *userRepository) VerifiedUsersCount() (*int64, error) {
	var totalVerifiedUsers int64

	err := r.db.Where("verification_status = ?", "verified").Model(&domain.User{}).Count(&totalVerifiedUsers).Error
	if err != nil {
		log.Printf("failed to count verified users")
		return nil,err
	}

	return &totalVerifiedUsers,nil
}


