package dto

type RegisterRequest struct {
	FirstName string `gorm:"size:100;not null" json:"first_name"`
	LastName  string `gorm:"size:100;not null" json:"last_name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	FirstName      string  `json:"first_name,omitempty"`
	LastName       string  `json:"last_name,omitempty"`
	Phone          string  `json:"phone,omitempty"`
	Gender         string  `json:"gender,omitempty"`
	DateOfBirth    DOB     `json:"date_of_birth,omitempty"` // "YYYY-MM-DD"
	ProfilePicture string  `json:"profile_picture,omitempty"`
	EnrollmentYear int     `json:"enrollment_year"`
	GraduationYear int     `json:"graduation_year"`
	Degree         string  `json:"degree"`
	Major          string  `json:"major"`
	RollNumber     *string `json:"roll_number"`
	CurrentCompany string  `json:"current_company,omitempty"`
	JobTitle       string  `json:"job_title,omitempty"`
	Industry       string  `json:"industry,omitempty"`
	LinkedInURL    string  `json:"linkedin_url,omitempty"`
	Website        string  `json:"website,omitempty"`
	City           string  `json:"city,omitempty"`
	State          string  `json:"state,omitempty"`
	Country        string  `json:"country,omitempty"`
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}
