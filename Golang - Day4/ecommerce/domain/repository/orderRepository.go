package repository

import (
	"ecommerce/domain/entity"

	"github.com/google/uuid"
)

type OrderRepository interface {
	GetOrderByID(uuid.UUID) (*entity.Order, error)
}
