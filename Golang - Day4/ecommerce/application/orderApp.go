package application

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"
)

type orderApp struct {
	orderRepo repository.OrderRepository
}

type OrderAppInterface interface {
	// GetOrderByID(uuid.UUID) (*entity.Order, error)
	AddOrder(*entity.Order) (*entity.Order, error)
	AddOrderDetail(*entity.OrderDetail) (*entity.OrderDetail, error)
}

var _ OrderAppInterface = &orderApp{}

// func (odr *orderApp) GetOrderByID(uuid uuid.UUID) (*entity.Order, error) {
// 	return odr.orderRepo.GetOrderByID(uuid)
// }

func (odr *orderApp) AddOrder(order *entity.Order) (*entity.Order, error) {
	return odr.orderRepo.AddOrder(order)
}

func (odr *orderApp) AddOrderDetail(orderDetail *entity.OrderDetail) (*entity.OrderDetail, error) {
	return odr.orderRepo.AddOrderDetail(orderDetail)
}
