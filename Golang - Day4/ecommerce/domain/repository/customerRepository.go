package repository

import (
	"ecommerce/domain/entity"
	"errors"
)

var ErrAddCustomer = errors.New("Couldn't add customer")

type CustomerRepository interface {
	AddCustomer(*entity.Customer) (*entity.Customer, error)
}
