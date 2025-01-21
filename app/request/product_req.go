package request

import "time"

type Product struct {
	Name              string    `json:"name" binding:"required" example:"apple"`
	Description       string    `json:"description" example:"apple"`
	Price             float64   `json:"price" binding:"required" example:"200"`
	Amount            int       `json:"amount" binding:"required" example:"10"`
	ProductCategoryId int       `json:"product_category_id" binding:"required" example:"1"`
	SaleOpenDate      time.Time `json:"sale_open_date" binding:"required" example:"2021-12-26T00:00:00Z"`
	SaleCloseDate     time.Time `json:"sale_close_date" binding:"required" example:"2021-12-26T00:00:00Z"`
}

type UpdateProduct struct {
	Name              string    `json:"name" example:"apple"`
	Description       string    `json:"description" example:"apple"`
	Price             float64   `json:"price" example:"200"`
	Amount            int       `json:"amount" example:"10"`
	ProductCategoryId int       `json:"product_category_id" example:"1"`
	SaleOpenDate      time.Time `json:"sale_open_date" example:"2021-12-26T00:00:00Z"`
	SaleCloseDate     time.Time `json:"sale_close_date" example:"2021-12-26T00:00:00Z"`
	UpdatedAt         time.Time `json:"updated_at"`
}
