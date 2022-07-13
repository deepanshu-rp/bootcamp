package repository

import "ecommerce/domain/entity"

type OrderRepository interface {
	// GetOrderByID(uuid.UUID) (*entity.Order, error)
	AddOrder(*entity.Order) (*entity.Order, error)
	AddOrderDetail(*entity.OrderDetail) (*entity.OrderDetail, error)
}
