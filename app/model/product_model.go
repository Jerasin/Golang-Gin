package model

import "time"

type Product struct {
	BaseModel
	Name              string `gorm:"unique;not null" json:"name" binding:"required"`
	Description       string
	Price             float64    `gorm:"not null"`
	Amount            int        `gorm:"not null"`
	SaleOpenDate      *time.Time `gorm:"column:sale_open_date" json:"saleOpenDate" example:"2021-12-26T00:00:00Z"`
	SaleCloseDate     *time.Time `gorm:"column:sale_close_date" json:"saleCloseDate" example:"2021-12-26T00:00:00Z"`
	ProductCategoryID uint       `gorm:"not null"`
	ImgUrl            string     `gorm:"column:img_url" json:"imgUrl,omitempty"`
}
