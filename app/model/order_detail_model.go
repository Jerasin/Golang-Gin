package model

type OrderDetail struct {
	BaseModel
	ProductID uint    `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Amount    int     `gorm:"not null"`
	OrderID   uint    `gorm:"not null"`
}

func (OrderDetail) TableName() string {
	return "orderDetails"
}
