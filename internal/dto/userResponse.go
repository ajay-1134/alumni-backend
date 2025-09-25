package dto

import "github.com/ajay-1134/alumni-backend/internal/domain"

type UserResponse struct {
	ID             uint   `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone,omitempty"`
	Gender         string `json:"gender,omitempty"`
	DateOfBirth    string `json:"date_of_birth,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`

	EnrollmentYear int     `json:"enrollment_year,omitempty"`
	GraduationYear int     `json:"graduation_year,omitempty"`
	Degree         string  `json:"degree,omitempty"`
	Major          string  `json:"major,omitempty"`
	RollNumber     *string `json:"roll_number,omitempty"`

	CurrentCompany string `json:"current_company,omitempty"`
	JobTitle       string `json:"job_title,omitempty"`
	Industry       string `json:"industry,omitempty"`
	LinkedinURL    string `json:"linkedin_url,omitempty"`
	Website        string `json:"website,omitempty"`

	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`

	Role string `json:"role"`
}

func ToUserResponse(user *domain.User) *UserResponse {
	dob := DOB(user.DateOfBirth)
	dobInBytes, _ := dob.MarshalJSON()
	return &UserResponse{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Phone:          user.Phone,
		Gender:         user.Gender,
		DateOfBirth:    string(dobInBytes),
		ProfilePicture: user.ProfilePicture,
		EnrollmentYear: user.EnrollmentYear,
		GraduationYear: user.GraduationYear,
		Degree:         user.Degree,
		Major:          user.Major,
		RollNumber:     user.RollNumber,
		CurrentCompany: user.CurrentCompany,
		JobTitle:       user.JobTitle,
		Industry:       user.Industry,
		LinkedinURL:    user.LinkedinURL,
		Website:        user.Website,
		City:           user.City,
		State:          user.State,
		Country:        user.Country,
		Role:           user.Role,
	}
}
