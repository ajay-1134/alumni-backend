package repository

import (
	"fmt"
	"log"

	"github.com/ajay-1134/alumni-backend/internal/domain"
	"github.com/ajay-1134/alumni-backend/internal/ports/repository"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *domain.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) GetAll() ([]*domain.Post, error) {
	var posts []*domain.Post

	if err := r.db.Find(&posts).Error; err != nil {
		log.Printf("error occured in getting posts")
		return nil, err
	}

	for _, post := range posts {
		imageUrl, err := generateSignedURL(post.ImageURL)
		fmt.Printf("imageUrl: %v\n", imageUrl)
		if err != nil {
			log.Printf("error occured in generating image url")
			return nil, err
		}

		post.ImageURL = imageUrl
	}

	return posts, nil
}

func (r *postRepository) GetAllWithUserId(userId uint) ([]*domain.Post, error) {
	var posts []*domain.Post

	err := r.db.Where("user_id = ?", userId).Find(&posts).Error
	if err != nil {
		log.Printf("error occured in getting posts for user %v", userId)
		return nil, err
	}

	for _, post := range posts {
		imageUrl, err := generateSignedURL(post.ImageURL)
		if err != nil {
			log.Printf("error occured in generating image url")
			return nil, err
		}

		post.ImageURL = imageUrl
	}

	return posts, nil
}

func (r *postRepository) GetPostById(id uint) (*domain.Post, error) {
	var post *domain.Post
	err := r.db.First(&post, id).Error
	if err != nil {
		log.Printf("error occured while getting post with id %v", id)
		return nil, err
	}

	imageUrl, err := generateSignedURL(post.ImageURL)
	if err != nil {
		log.Printf("error occured in generating image url")
		return nil, err
	}

	post.ImageURL = imageUrl

	return post, nil
}

func (r *postRepository) Update(post *domain.Post, updates *map[string]interface{}) error {
	if err := r.db.Model(&post).Updates(updates).Error; err != nil {
		log.Printf("error occured in updating post")
		return err
	}

	return nil
}

func (r *postRepository) Delete(postId uint) error {
	return r.db.Delete(&domain.Post{}, postId).Error
}
