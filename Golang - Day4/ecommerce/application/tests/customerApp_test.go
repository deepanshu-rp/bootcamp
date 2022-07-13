package application_test

import (
	"ecommerce/domain/entity"
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
	testVal, err := mockApp.GetCustomerByID(customer1.CustomerId)
	if err == nil {
		assert.Equal(t, &customer1, testVal)
	}
}

func TestAddCustomer(t *testing.T) {
	mockApp := mocks.NewCustomerAppInterface(t)

	// Error??
	// mockApp.On("AddCustomer", mock.AnythingOfType("*entity.Customer")).Return(
	// 	func(c *entity.Customer) (*entity.Customer, error) {
	// 		return c, nil
	// 	},
	// )

	mockApp.On("AddCustomer", mock.AnythingOfType("*entity.Customer")).Return(&customer1, nil)
	testVal, err := mockApp.AddCustomer(&customer2)
	if err == nil {
		assert.Equal(t, &customer1, testVal)
	}
}
