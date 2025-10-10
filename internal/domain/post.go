package domain

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"` // foreign key
	Text      string    `gorm:"type:text;not null" json:"text"`
	ImageURL  string    `json:"image_url"` // store the uploaded image link
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relationship
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
