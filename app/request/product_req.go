package request

import "time"

type Product struct {
	Name              string    `json:"name" binding:"required" example:"apple"`
	Description       string    `json:"description" example:"apple"`
	Price             float64   `json:"price" binding:"required" example:"200"`
	Amount            int       `json:"amount" binding:"required" example:"10"`
	ProductCategoryId int       `json:"productCategoryId" binding:"required" example:"1"`
	SaleOpenDate      time.Time `json:"saleOpenDate"  example:"2021-12-26T00:00:00Z"`
	SaleCloseDate     time.Time `json:"saleCloseDate" example:"2021-12-26T00:00:00Z"`
	ImgUrl            string    `json:"imgUrl"`
}

type UpdateProduct struct {
	Name              string    `json:"name" example:"apple"`
	Description       string    `json:"description" example:"apple"`
	Price             float64   `json:"price" example:"200"`
	Amount            int       `json:"amount" example:"10"`
	ProductCategoryId int       `json:"product_category_id" example:"1"`
	SaleOpenDate      time.Time `json:"saleOpenDate" example:"2021-12-26T00:00:00Z"`
	SaleCloseDate     time.Time `json:"saleCloseDate" example:"2021-12-26T00:00:00Z"`
	UpdatedAt         time.Time `json:"updatedAt"`
	ImgUrl            string    `json:"imgUrl"`
}
