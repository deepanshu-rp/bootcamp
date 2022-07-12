package persistence

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"

	"github.com/google/uuid"
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

func (pgr *ProductGormRepo) GetProductByID(uuid uuid.UUID) (*entity.Product, error) {
	var product entity.Product

	if err := pgr.db.Where("product_id = ?", uuid).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (pgr *ProductGormRepo) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	if err := pgr.db.Model(&entity.Product{}).Where("product_id = ?", product.ProductId).Update("product_quantity", product.ProductQuantity).Error; err != nil {
		return nil, err
	}
	return product, nil
}
