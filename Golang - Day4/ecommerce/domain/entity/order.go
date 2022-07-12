package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Order struct {
	OrderId     uuid.UUID `json:"order_id" gorm:"order_id"`
	OrderStatus string    `json:"order_status" gorm:"customer_id"`
	CustomerId  uuid.UUID `json:"customer_id" gorm:"customer_id"`
}

type Orders []Order

type OrderDetail struct {
	OrderId   uuid.UUID `json:"order_id" gorm:"order_id"`
	ProductId uuid.UUID `json:"product_id" gorm:"product_id"`
	Quantity  int       `json:"quantity" gorm:"quantity"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}

func (o *Order) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New().String()
	return scope.SetColumn("order_id", id)

}

func (o *OrderDetail) OnCreate(scope *gorm.Scope) error {
	time := time.Now()
	return scope.SetColumn("created_at", time)
}

func (o Order) TableName() string {
	return "order"
}
