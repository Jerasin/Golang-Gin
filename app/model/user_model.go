package model

type User struct {
	BaseModel
	Username   string   `gorm:"unique;not null" json:"username,omitempty"`
	Password   string   `gorm:"not null" json:"password,omitempty"`
	Fullname   string   `gorm:"unique;not null" json:"fullname,omitempty"`
	Avatar     string   `json:"avatar,omitempty"`
	Email      string   `gorm:"unique;not null"`
	Order      []Order  `gorm:"foreignKey:CreatedBy;references:ID" json:"ordes,omitempty"`
	RoleInfoID uint     `gorm:"not null"`
	Wallets    []Wallet `json:"allets,omitempty"`
}
