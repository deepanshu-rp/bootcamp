package application

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"

	"github.com/google/uuid"
)

type productApp struct {
	product repository.ProductRepository
}

type ProductAppInterface interface {
	AddProduct(*entity.Product) (*entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	GetProductByID(uuid.UUID) (*entity.Product, error)
}

var _ ProductAppInterface = &productApp{}

func (p *productApp) AddProduct(product *entity.Product) (*entity.Product, error) {
	return p.product.AddProduct(product)
}

func (p *productApp) GetAllProducts() ([]entity.Product, error) {
	return p.product.GetAllProducts()
}
func (p *productApp) GetProductByID(uuid uuid.UUID) (*entity.Product, error) {
	return p.product.GetProductByID(uuid)
}
