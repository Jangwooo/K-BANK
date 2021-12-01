package DAO

type SimplePwd struct {
	UserID string `gorm:"primaryKey" json:"user_id,omitempty"`
	Pwd    string `json:"pwd,omitempty"`
}
