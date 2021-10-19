package model

type BankInfo struct {
	ID       string `gorm:"primaryKey"`
	BankName string `gorm:"unique"`

	BankLogo BankLogo `gorm:"foreignKey: BankID"`

	AnotherAccount  []AnotherAccount  `gorm:"foreignKey: BankID"`
	CheckingAccount []CheckingAccount `gorm:"foreignKey: BankID"`
}
