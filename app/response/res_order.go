package response

import "github.com/Jerasin/app/model"

type Order struct {
	model.BaseModel
	TotalPrice  float64 `json:"totalPrice" binding:"required"`
	TotalAmount int     `json:"totalAmount" binding:"required"`
}

type OrderPagination struct {
	PaginationResponse
	Data []Order `json:"data"`
}
