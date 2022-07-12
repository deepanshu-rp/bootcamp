package entity

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	ProductId       uuid.UUID `json:"product_id" gorm:"product_id"`
	ProductName     string    `json:"product_name" gorm:"product_name"`
	ProductPrice    int       `json:"product_price" gorm:"product_price"`
	ProductQuantity int       `json:"product_quantity" gorm:"product_quantity"`
	RetailerId      uuid.UUID `json:"retailer_id" gorm:"retailer_id"`
}

type ProductPatch struct {
	ProductPrice    int `json:"product_price"`
	ProductQuantity int `json:"product_quantity"`
}

func (p Product) TableName() string {
	return "product"
}
func (product *Product) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New().String()
	return scope.SetColumn("product_id", id)
}
