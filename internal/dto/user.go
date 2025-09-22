package dto

type RegisterRequest struct{
	// Personal Info
	FirstName      string    `gorm:"size:100;not null" json:"first_name"`
	LastName       string    `gorm:"size:100;not null" json:"last_name"`
	Email          string    `gorm:"unique;not null" json:"email"`
	Phone          string    `gorm:"size:15" json:"phone"`
	Gender         string    `gorm:"size:10" json:"gender"`
	DateOfBirth    string `gorm:"type:date" json:"date_of_birth"`
	ProfilePicture string    `json:"profile_picture"`

	// Academic Info
	EnrollmentYear int    `json:"enrollment_year"`
	GraduationYear int    `json:"graduation_year"`
	Degree         string `gorm:"size:100" json:"degree"`
	Major          string `gorm:"size:100" json:"major"`
	RollNumber     string `gorm:"size:50;unique" json:"roll_number"`

	// Professional Info
	CurrentCompany string `gorm:"size:150" json:"current_company"`
	JobTitle       string `gorm:"size:100" json:"job_title"`
	Industry       string `gorm:"size:100" json:"industry"`
	LinkedInURL    string `json:"linkedin_url"`
	Website        string `json:"website"`

	// Location
	City    string `gorm:"size:100" json:"city"`
	State   string `gorm:"size:100" json:"state"`
	Country string `gorm:"size:100" json:"country"`

	Password  string `json:"password" binding:"required"`
	AuthProvider string `json:"auth_provider"`
	AuthID       string `json:"auth_id"`
}

type UpdateUserRequest struct {
    FirstName      string `json:"first_name,omitempty"`
    LastName       string `json:"last_name,omitempty"`
    Phone          string `json:"phone,omitempty"`
    Gender         string `json:"gender,omitempty"`
    DateOfBirth    string `json:"date_of_birth,omitempty"` // "YYYY-MM-DD"
    ProfilePicture string `json:"profile_picture,omitempty"`

    Degree         string `json:"degree,omitempty"`
    Major          string `json:"major,omitempty"`

    CurrentCompany string `json:"current_company,omitempty"`
    JobTitle       string `json:"job_title,omitempty"`
    Industry       string `json:"industry,omitempty"`
    LinkedInURL    string `json:"linkedin_url,omitempty"`
    Website        string `json:"website,omitempty"`

    City    string `json:"city,omitempty"`
    State   string `json:"state,omitempty"`
    Country string `json:"country,omitempty"`
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

