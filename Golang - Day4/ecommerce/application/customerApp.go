package application

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"
)

type customerApp struct {
	cust repository.CustomerRepository
}

type CustomerAppInterface interface {
	AddCustomer(*entity.Customer) (*entity.Customer, error)
}

var _ CustomerAppInterface = &customerApp{}

func (c *customerApp) AddCustomer(customer *entity.Customer) (*entity.Customer, error) {
	return c.cust.AddCustomer(customer)
}
