package persistence

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
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

func (rgr *RetailerGormRepo) GetRetailerByID(uuid uuid.UUID) (*entity.Retailer, error) {
	var retailer entity.Retailer

	if err := rgr.db.Where("retailer_id = ?", uuid).Find(&retailer).Error; err != nil {
		return nil, err
	}
	return &retailer, nil
}
