package entity

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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

var uuidRegex = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"

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

func (p Product) ValidateInput() error {
	err := validation.ValidateStruct(&p,
		validation.Field(&p.ProductId, validation.Required, validation.Match(regexp.MustCompile(uuidRegex))),
		validation.Field(&p.ProductName, validation.Required, validation.Length(1, 20)),
		validation.Field(&p.ProductPrice, validation.Required, validation.Min(0)),
		validation.Field(&p.ProductQuantity, validation.Required, validation.Min(0)),
		validation.Field(&p.RetailerId, validation.Required, validation.Match(regexp.MustCompile(uuidRegex))),
	)
	if err != nil {
		return err
	}
	return nil
}
