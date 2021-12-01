package DAO

type AnotherAccount struct {
	ID     string `gorm:"primaryKey" json:"id,omitempty"`
	UserID string `json:"user_id,omitempty"`
	BankID string `json:"bank_id,omitempty"`
}
