package model

type Wallet struct {
	BaseModel
	Name   string  `gorm:"unique;not null"`
	Token  string  `gorm:"unique;not null"`
	Uuid   string  `gorm:"unique;not null"`
	UserID uint    `gorm:"not null" json:"user_id"`
	Value  float64 `gorm:"not null"`
	User   User
}
