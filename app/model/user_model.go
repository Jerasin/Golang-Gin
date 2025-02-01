package model

type User struct {
	BaseModel
	Username   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	Fullname   string `gorm:"unique;not null"`
	Avatar     string
	Email      string  `gorm:"unique;not null"`
	Order      []Order `gorm:"foreignKey:CreatedBy;references:ID"`
	RoleInfoID uint    `gorm:"not null"`
	Wallets    []Wallet
}
