package model

type SimplePwd struct {
	UserID string `gorm:"primaryKey"`
	Pwd    string
}
