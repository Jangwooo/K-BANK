package model

import (
	"database/sql"
)

type User struct {
	ID       string `gorm:"primaryKey"`
	Password string
	SSN      string `gorm:"unique"`
	Name     string
	nickName sql.NullString
	agree    string
	UserType string

	ProfilePic ProfilePic
	SimplePwd  SimplePwd

	CheckingAccount []CheckingAccount `gorm:"foreignKey:UserID"`
}
