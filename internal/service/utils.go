package service

import (
	"github.com/ajay-1134/alumni-backend/internal/domain"
	"github.com/ajay-1134/alumni-backend/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

// func updateRequestToUser(req *dto.UpdateUserRequest) *domain.User {

// 	user := domain.User{
// 		FirstName:      req.FirstName,
// 		LastName:       req.LastName,
// 		Phone:          req.Phone,
// 		Gender:         req.Gender,
// 		DateOfBirth:    *convertDOB(req.DateOfBirth),
// 		ProfilePicture: req.ProfilePicture,
// 		Degree:         req.Degree,
// 		Major:          req.Major,
// 		CurrentCompany: req.CurrentCompany,
// 		JobTitle:       req.JobTitle,
// 		Industry:       req.Industry,
// 		LinkedInURL:    req.LinkedInURL,
// 		Website:        req.Website,
// 		City:           req.City,
// 		State:          req.City,
// 		Country:        req.Country,
// 	}

// 	return &user
// }

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
