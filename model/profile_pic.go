package model

type ProfilePic struct {
	UserID string `gorm:"primaryKey" json:"user_id,omitempty"`
	Path   string `json:"path,omitempty"`
}
