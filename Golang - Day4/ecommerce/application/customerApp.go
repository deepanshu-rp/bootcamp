package application

import (
	"ecommerce/domain/entity"
	"ecommerce/domain/repository"

	"github.com/google/uuid"
)

type customerApp struct {
	cust repository.CustomerRepository
}

type CustomerAppInterface interface {
	AddCustomer(*entity.Customer) (*entity.Customer, error)
	GetCustomerByID(uuid.UUID) (*entity.Customer, error)
}

var _ CustomerAppInterface = &customerApp{}

func (c *customerApp) AddCustomer(customer *entity.Customer) (*entity.Customer, error) {
	return c.cust.AddCustomer(customer)
}

func (c *customerApp) GetCustomerByID(uuid uuid.UUID) (*entity.Customer, error) {
	return c.cust.GetCustomerByID(uuid)
}
