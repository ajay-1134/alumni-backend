package service

import (
	"github.com/ajay-1134/alumni-backend/internal/domain"
	"github.com/ajay-1134/alumni-backend/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

func registerRequestToUser(req *dto.RegisterRequest) *domain.User {

	user := domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	return &user
}

func googleUserToUser(googleUser *dto.GoogleUser) *domain.User {
	return &domain.User{
		FirstName:      googleUser.GivenName,
		LastName:       googleUser.FamilyName,
		Email:          googleUser.Email,
		ProfilePicture: googleUser.Picture,
		AuthProvider:   "Google",
		AuthID:         googleUser.ID,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func buildUpdates(req *dto.UpdateUserRequest) *map[string]interface{} {
	updates := make(map[string]interface{})

	if req.FirstName != "" {
		updates["first_name"] = req.FirstName
	}
	if req.LastName != "" {
		updates["last_name"] = req.LastName
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.Gender != "" {
		updates["gender"] = req.Gender
	}
	if !req.DateOfBirth.ToTime().IsZero() {
		updates["date_of_birth"] = req.DateOfBirth.ToTime()
	}
	if req.ProfilePicture != "" {
		updates["profile_picture"] = req.ProfilePicture
	}
	if req.RollNumber != nil {
		updates["roll_number"] = req.RollNumber
	}
	if req.Degree != "" {
		updates["degree"] = req.Degree
	}
	if req.Major != "" {
		updates["major"] = req.Major
	}
	if req.EnrollmentYear != 0 {
		updates["enrollment_year"] = req.EnrollmentYear
	}
	if req.GraduationYear != 0 {
		updates["graduation_year"] = req.GraduationYear
	}
	if req.CurrentCompany != "" {
		updates["current_company"] = req.CurrentCompany
	}
	if req.JobTitle != "" {
		updates["job_title"] = req.JobTitle
	}
	if req.Industry != "" {
		updates["industry"] = req.Industry
	}
	if req.LinkedInURL != "" {
		updates["linkedin_url"] = req.LinkedInURL
	}
	if req.Website != "" {
		updates["website"] = req.Website
	}
	if req.City != "" {
		updates["city"] = req.City
	}
	if req.State != "" {
		updates["state"] = req.State
	}
	if req.Country != "" {
		updates["country"] = req.Country
	}

	return &updates
}

func dtoToPost(postDto *dto.PostDto) *domain.Post {
	return &domain.Post{
		ID:       postDto.PostID,
		UserID:   postDto.UserID,
		Text:     postDto.Text,
		ImageURL: postDto.ImageURL,
		User:     domain.User{},
	}
}

func postToDto(post *domain.Post) *dto.PostDto {
	return &dto.PostDto{
		PostID:   post.ID,
		ImageURL: post.ImageURL,
		Text:     post.Text,
		UserID:   post.UserID,
	}
}

func buildUpdatedPost(post *dto.PostDto) *map[string]interface{} {
	updatedPost := make(map[string]interface{})

	if post.ImageURL != "" {
		updatedPost["image_url"] = post.ImageURL
	}
	if post.Text != "" {
		updatedPost["text"] = post.Text
	}

	return &updatedPost
}
