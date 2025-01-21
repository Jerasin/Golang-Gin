package response

type Order struct {
	Id          int     `json:"id"`
	TotalPrice  float64 `json:"total_price" binding:"required"`
	TotalAmount int     `json:"total_amount" binding:"required"`
}

type OrderPagination struct {
	PaginationResponse
	Data []Order `json:"data"`
}
