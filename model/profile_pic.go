package model

type ProfilePic struct {
	UserID string `gorm:"primaryKey"`
	Path   string
}
