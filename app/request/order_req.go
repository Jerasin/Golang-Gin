package request

type OrderRequest struct {
	Orders   []OrderItem `json:"orders" binding:"required"`
	WalletID uint        `json:"wallet_id" binding:"required" example:"10"`
}

type OrderItem struct {
	ProductId int `json:"product_id" binding:"required" example:"1"`
	Amount    int `json:"amount" binding:"required" example:"10"`
}
