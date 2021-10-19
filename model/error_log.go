package model

type ErrorLog struct {
	ID      string `gorm:"primaryKey"`
	TradeID string `gorm:"varchar(191)"`
	Detail  string
}
