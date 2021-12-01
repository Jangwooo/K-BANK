package DAO

type ErrorLog struct {
	ID      string `gorm:"primaryKey" json:"id,omitempty"`
	TradeID string `gorm:"varchar(191)" json:"trade_id,omitempty"`
	Detail  string `json:"detail,omitempty"`
}
