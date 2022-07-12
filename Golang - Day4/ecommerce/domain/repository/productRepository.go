package repository

import "ecommerce/domain/entity"

type ProductRepository interface {
	AddProduct(*entity.Product) (*entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
}
