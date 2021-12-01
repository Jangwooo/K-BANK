package DAO

import "time"

type History struct {
	ID        int       `gorm:"primaryKey; autoIncrement" json:"id,omitempty"`
	TradeType string    `json:"trade_type,omitempty"`
	AccountID string    `gorm:"varchar(191)" json:"account_id,omitempty"`
	Client    string    `json:"client,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Amount    int       `json:"amount,omitempty"`
	Balance   int       `json:"balance,omitempty"`
	State     string    `json:"state,omitempty"`

	ErrorLog *[]ErrorLog `gorm:"foreignKey: TradeID" json:"error_log,omitempty"`
}
