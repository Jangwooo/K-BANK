package model

import (
	"database/sql"
	"time"
)

type CheckingAccount struct {
	ID              string         `gorm:"primaryKey" json:"id,omitempty"`
	BankID          string         `json:"bank_id,omitempty"`
	UserID          string         `json:"user_id,omitempty"`
	Password        string         `json:"password,omitempty"`
	AccountNickname sql.NullString `json:"account_nickname"`
	Balance         int            `json:"balance,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	State           string         `json:"state,omitempty"`
	Limit           int            `json:"limit,omitempty"`

	History *[]History `gorm:"foreignKey:AccountID" json:"history,omitempty"`
}
