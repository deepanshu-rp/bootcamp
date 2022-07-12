package repository

import (
	"ecommerce/domain/entity"

	uuid "github.com/satori/go.uuid"
)

type RetailerRepository interface {
	AddRetailer(*entity.Retailer) (*entity.Retailer, error)
	GetRetailerByID(uuid.UUID) (*entity.Retailer, error)
}
