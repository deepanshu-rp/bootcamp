package entity

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Order struct {
	OrderId    uuid.UUID `gorm:"order_id"`
	CustomerId uuid.UUID `gorm:"customer_id"`
}

type OrderDetail struct {
	OrderId         uuid.UUID `json:"order_id" gorm:"order_id"`
	ProductId       uuid.UUID `json:"product_id" gorm:"product_id"`
	QuantityOrdered int       `json:"quantity_ordered" gorm:"quantity_ordered"`
	OrderStatus     string    `json:"order_status" gorm:"order_status"`
}

// For parsing order json
type OrderRecieved struct {
	OrderId    uuid.UUID   `json:"order_id"`
	CustomerId uuid.UUID   `json:"customer_id"`
	Specs      []OrderSpec `json:"order_detail"`
}

type OrderSpec struct {
	ProductId       uuid.UUID `json:"product_id"`
	ProductQuantity int       `json:"product_quantity"`
}

func (o *Order) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New().String()
	return scope.SetColumn("order_id", id)

}

func (o Order) TableName() string {
	return "order"
}
func (od OrderDetail) TableName() string {
	return "order_detail"
}
