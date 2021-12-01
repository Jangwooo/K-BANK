package DAO

import (
	"database/sql"
)

type User struct {
	ID          string         `gorm:"primaryKey" json:"id,omitempty"`
	Password    string         `json:"password,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	SSN         string         `gorm:"unique" json:"ssn,omitempty"`
	Name        string         `json:"name,omitempty"`
	NickName    sql.NullString `json:"nick_name"`
	Agree       string         `json:"agree,omitempty"`
	UserType    string         `json:"user_type,omitempty"`

	ProfilePic *ProfilePic `json:"profile_pic"`
	SimplePwd  *SimplePwd  `json:"simple_pwd"`

	CheckingAccount *[]CheckingAccount `gorm:"foreignKey:UserID" json:"checking_account,omitempty"`
}
