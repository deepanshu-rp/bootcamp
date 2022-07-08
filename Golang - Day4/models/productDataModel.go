package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	Id          string `json:"id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
}

func (product *Product) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	scope.SetColumn("Id", id.String())
	return nil
}
