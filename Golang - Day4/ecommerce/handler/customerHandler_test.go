package handler_test

import (
	"ecommerce/domain/entity"
	"ecommerce/handler"
	mocks "ecommerce/mocks/application"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	customer1 = entity.Customer{
		CustomerId:      uuid.MustParse("e3547962-01aa-4418-8b03-21103a9cedc1"),
		CustomerName:    "Test",
		CustomerAddress: "Home",
	}
	customer2 = entity.Customer{
		CustomerId:      uuid.MustParse("e9ea26fd-12ab-44e7-baf8-77890d7d27ac"),
		CustomerName:    "Test",
		CustomerAddress: "Home",
	}
)

func TestGetCustomerByID(t *testing.T) {
	mockApp := mocks.NewCustomerAppInterface(t)
	mockApp.On("GetCustomerByID", mock.AnythingOfType("uuid.UUID")).Return(&customer1, nil)
	testApp := handler.NewCustomerService(mockApp)

	customers := []entity.Customer{customer1, customer1}
	for _, cust := range customers {
		result, _ := testApp.Cust.GetCustomerByID(cust.CustomerId)
		assert.Equal(t, result, &cust)
	}

}

func TestAddCustomer(t *testing.T) {
	mockApp := mocks.NewCustomerAppInterface(t)
	customers := []entity.Customer{customer1, customer2}

	mockApp.On("AddCustomer", mock.AnythingOfType("*entity.Customer")).Return(
		func(c *entity.Customer) *entity.Customer {
			return c
		},
		func(c *entity.Customer) error {
			return nil
		},
	)
	testApp := handler.NewCustomerService(mockApp)
	for _, cust := range customers {
		result, _ := testApp.Cust.AddCustomer(&cust)
		assert.Equal(t, result, &cust)
	}

}
