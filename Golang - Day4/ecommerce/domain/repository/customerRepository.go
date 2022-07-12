package repository

import (
	"ecommerce/domain/entity"
	"errors"

	uuid "github.com/satori/go.uuid"
)

var ErrAddCustomer = errors.New("Couldn't add customer")

type CustomerRepository interface {
	AddCustomer(*entity.Customer) (*entity.Customer, error)
	GetCustomerByID(uuid.UUID) (*entity.Customer, error)
}
