package repository

import "ecommerce/domain/entity"

type RetailerRepository interface {
	AddRetailer(*entity.Retailer) (*entity.Retailer, error)
}
