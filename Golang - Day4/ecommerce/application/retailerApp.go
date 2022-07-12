package application

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"

	uuid "github.com/satori/go.uuid"
)

type retailerApp struct {
	retail repository.RetailerRepository
}

type RetailerAppInterface interface {
	AddRetailer(*entity.Retailer) (*entity.Retailer, error)
	GetRetailerByID(uuid.UUID) (*entity.Retailer, error)
}
