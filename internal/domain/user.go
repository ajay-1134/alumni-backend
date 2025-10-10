package domain

import (
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Personal Info
	FirstName      string    `gorm:"size:100;not null" json:"first_name"`
	LastName       string    `gorm:"size:100;not null" json:"last_name"`
	Email          string    `gorm:"unique;not null" json:"email"`
	Phone          string    `gorm:"size:15" json:"phone"`
	Gender         string    `gorm:"size:10" json:"gender"` // Male, Female, Other
	DateOfBirth    time.Time `gorm:"type:date" json:"date_of_birth"`
	ProfilePicture string    `json:"profile_picture"` // store URL

	// Academic Info
	EnrollmentYear int     `json:"enrollment_year"`
	GraduationYear int     `json:"graduation_year"`
	Degree         string  `gorm:"size:100" json:"degree"`
	Major          string  `gorm:"size:100" json:"major"`
	RollNumber     *string `gorm:"size:50;unique" json:"roll_number"`

	// Professional Info
	CurrentCompany string `gorm:"size:150" json:"current_company"`
	JobTitle       string `gorm:"size:100" json:"job_title"`
	Industry       string `gorm:"size:100" json:"industry"`
	LinkedinURL    string `json:"linkedin_url"`
	Website        string `json:"website"`

	// Location
	City    string `gorm:"size:100" json:"city"`
	State   string `gorm:"size:100" json:"state"`
	Country string `gorm:"size:100" json:"country"`

	// Metadata
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	PasswordHash string `json:"-"`
	AuthProvider string `json:"auth_provider"`
	AuthID       string `json:"auth_id"`

	Role  string        `gorm:"size:20;default:user" json:"role"`
	Posts []Post `json:"posts"`
}
