package persistence

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"

	"github.com/jinzhu/gorm"
)

type RetailerGormRepo struct {
	db *gorm.DB
}

func NewRetailerGormRepo(db *gorm.DB) *RetailerGormRepo {
	return &RetailerGormRepo{db: db}
}

var _ repository.RetailerRepository = &RetailerGormRepo{}

func (rgr *RetailerGormRepo) AddRetailer(retailer *entity.Retailer) (*entity.Retailer, error) {
	if err := rgr.db.Create(retailer).Error; err != nil {
		return nil, err
	}
	return retailer, nil
}
