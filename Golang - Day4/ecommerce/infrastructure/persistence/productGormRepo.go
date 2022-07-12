package persistence

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"

	"github.com/jinzhu/gorm"
)

type ProductGormRepo struct {
	db *gorm.DB
}

func NewProductGormRepo(db *gorm.DB) *ProductGormRepo {
	return &ProductGormRepo{db: db}
}

var _ repository.ProductRepository = &ProductGormRepo{}

func (pgr *ProductGormRepo) AddProduct(product *entity.Product) (*entity.Product, error) {
	if err := pgr.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (pgr *ProductGormRepo) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	err := pgr.db.Debug().Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
