package repository

import (
	"ecommerce/domain/entity"

	"github.com/google/uuid"
)

type ProductRepository interface {
	AddProduct(*entity.Product) (*entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	GetProductByID(uuid.UUID) (*entity.Product, error)
}
