package application

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"

	"github.com/google/uuid"
)

type retailerApp struct {
	retail repository.RetailerRepository
}

type RetailerAppInterface interface {
	AddRetailer(*entity.Retailer) (*entity.Retailer, error)
	GetRetailerByID(uuid.UUID) (*entity.Retailer, error)
}
