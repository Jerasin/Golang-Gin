package request

type OrderRequest struct {
	Orders   []OrderItem `json:"orders" binding:"required"`
	WalletID uint        `json:"walletId" binding:"required" example:"10"`
}

type OrderItem struct {
	ProductId int `json:"productId" binding:"required" example:"1"`
	Amount    int `json:"amount" binding:"required" example:"10"`
}
