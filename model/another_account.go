package model

type AnotherAccount struct {
	ID     string `gorm:"primaryKey"`
	UserID string
	BankID string
}
