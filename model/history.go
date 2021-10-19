package model

import "time"

type History struct {
	ID        string `gorm:"primaryKey"`
	TradeType string
	AccountID string `gorm:"varchar(191)"`
	Client    string
	CreatedAt time.Time
	Amount    int
	Balance   int
	State     string

	ErrorLog []ErrorLog `gorm:"foreignKey: TradeID"`
}
