package dto

type PostDto struct {
	PostID uint `json:"id,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	Text string `json:"text,omitempty"`
	UserID uint `json:"user_id,omitempty"`
}