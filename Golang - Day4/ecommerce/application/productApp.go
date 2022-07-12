package application

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"
)

type productApp struct {
	product repository.ProductRepository
}

type ProductAppInterface interface {
	AddProduct(*entity.Product) (*entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
}

var _ ProductAppInterface = &productApp{}

func (p *productApp) AddProduct(product *entity.Product) (*entity.Product, error) {
	return p.product.AddProduct(product)
}

func (p *productApp) GetAllProducts() ([]entity.Product, error) {
	return p.product.GetAllProducts()
}
