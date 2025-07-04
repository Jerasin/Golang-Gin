package response

import "time"

type Product struct {
	Id            int        `json:"id"`
	Name          string     `json:"name" binding:"required" example:"apple"`
	Description   string     `json:"description" example:"apple"`
	Price         float64    `json:"price" binding:"required" example:"200"`
	Amount        int        `json:"amount" binding:"required" example:"10"`
	ImgUrl        string     `gorm:"column:img_url" json:"imgUrl"`
	SaleOpenDate  *time.Time `gorm:"column:sale_open_date" json:"saleOpenDate" binding:"required" example:"2021-12-26T00:00:00Z"`
	SaleCloseDate *time.Time `gorm:"column:sale_close_date" json:"saleCloseDate" binding:"required" example:"2021-12-26T00:00:00Z"`
}

type ProductPagination struct {
	PaginationResponse
	Data []Product `json:"data"`
}
