package model

import (
	"database/sql"
	"time"
)

type CheckingAccount struct {
	ID              string `gorm:"primaryKey"`
	BankID          string
	UserID          string
	Password        string
	AccountNickname sql.NullString
	Balance         int
	CreatedAt       time.Time
	State           string
	Limit           int

	History []History `gorm:"foreignKey:AccountID"`
}
