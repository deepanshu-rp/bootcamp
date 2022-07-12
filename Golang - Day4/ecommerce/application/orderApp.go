package application

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"

	"github.com/google/uuid"
)

type orderApp struct {
	orderRepo repository.OrderRepository
}

type OrderAppInterface interface {
	GetOrderByID(uuid.UUID) (*entity.Order, error)
}

var _ OrderAppInterface = &orderApp{}

func (odr *orderApp) GetOrderByID(uuid uuid.UUID) (*entity.Order, error) {
	return odr.orderRepo.GetOrderByID(uuid)
}
