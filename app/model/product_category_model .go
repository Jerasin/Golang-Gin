package model

type ProductCategory struct {
	BaseModel
	Name        string `gorm:"unique;not null"`
	Description string
	Products    []Product
}

func (ProductCategory) TableName() string {
	return "productCategories"
}
