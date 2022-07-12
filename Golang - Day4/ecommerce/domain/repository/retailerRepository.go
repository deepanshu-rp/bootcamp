package repository

import (
	"ecommerce/domain/entity"

	"github.com/google/uuid"
)

type RetailerRepository interface {
	AddRetailer(*entity.Retailer) (*entity.Retailer, error)
	GetRetailerByID(uuid.UUID) (*entity.Retailer, error)
}
