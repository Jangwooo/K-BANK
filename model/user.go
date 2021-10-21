package model

import (
	"database/sql"
)

type User struct {
	ID          string `gorm:"primaryKey"`
	Password    string
	PhoneNumber string
	SSN         string `gorm:"unique"`
	Name        string
	NickName    sql.NullString
	Agree       string
	UserType    string

	ProfilePic ProfilePic
	SimplePwd  SimplePwd

	CheckingAccount []CheckingAccount `gorm:"foreignKey:UserID"`
}
