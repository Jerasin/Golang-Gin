package model

import "time"

type Product struct {
	BaseModel
	Name              string `gorm:"unique;not null" json:"name" binding:"required"`
	Description       string
	Price             float64 `gorm:"not null"`
	Amount            int     `gorm:"not null"`
	SaleOpenDate      *time.Time
	SaleCloseDate     *time.Time
	ProductCategoryID uint   `gorm:"not null"`
	ImgUrl            string `gorm:"default:null"`
}
