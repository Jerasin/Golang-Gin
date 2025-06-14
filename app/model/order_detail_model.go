package model

type OrderDetail struct {
	BaseModel
	ProductID   uint    `gorm:"not null" json:"productId"`
	ProductName string  `gorm:"not null" json:"productName"`
	Price       float64 `gorm:"not null" json:"price"`
	Amount      int     `gorm:"not null" json:"amount"`
	OrderID     uint    `gorm:"not null" json:"orderId"`
}

func (OrderDetail) TableName() string {
	return "orderDetails"
}
