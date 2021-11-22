package model

type BankInfo struct {
	ID       string `gorm:"primaryKey" json:"id,omitempty"`
	BankName string `gorm:"unique" json:"bank_name,omitempty"`

	BankLogo BankLogo `gorm:"foreignKey: BankID" json:"bank_logo"`

	AnotherAccount  *[]AnotherAccount  `gorm:"foreignKey: BankID" json:"another_account,omitempty"`
	CheckingAccount *[]CheckingAccount `gorm:"foreignKey: BankID" json:"checking_account,omitempty"`
}
