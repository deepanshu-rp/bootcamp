package repository

import (
	"ecommerce/domain/entity"

	"github.com/google/uuid"
)

type OrderRepository interface {
	GetOrderByID(uuid.UUID) ([]entity.OrderDetail, error)
	AddOrder(*entity.Order) (*entity.Order, error)
	AddOrderDetail(*entity.OrderDetail) (*entity.OrderDetail, error)
}
