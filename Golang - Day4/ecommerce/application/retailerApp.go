package application

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"
)

type retailerApp struct {
	retail repository.RetailerRepository
}

type RetailerAppInterface interface {
	AddRetailer(*entity.Retailer) (*entity.Retailer, error)
}
