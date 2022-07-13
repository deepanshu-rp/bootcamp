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
	product1 = entity.Product{
		ProductId:       uuid.MustParse("e9ea26fd-12ab-44e7-baf8-77890d7d27ac"),
		RetailerId:      uuid.MustParse("e3547962-01aa-4418-8b03-21103a9cedc1"),
		ProductName:     "Product 1",
		ProductQuantity: 10,
		ProductPrice:    4,
	}
	product2 = entity.Product{
		ProductId:       uuid.MustParse("e3547962-01aa-4418-8b03-21103a9cedc1"),
		RetailerId:      uuid.MustParse("e9ea26fd-12ab-44e7-baf8-77890d7d27ac"),
		ProductName:     "Product 2",
		ProductQuantity: 200,
		ProductPrice:    50,
	}
)

func TestUpdateProduct(t *testing.T) {
	mockApp := mocks.NewProductAppInterface(t)
	mockApp.On("UpdateProduct", mock.AnythingOfType("*entity.Product")).Return(&product1, nil)
	testVal, err := mockApp.UpdateProduct(&product2)
	if err == nil {
		assert.Equal(t, &product1, testVal)
	}
	mockApp.AssertExpectations(t)
}

func TestGetAllProducts(t *testing.T) {
	mockApp := mocks.NewProductAppInterface(t)
	productsMock := []entity.Product{product1, product2}
	productsExpected := []entity.Product{product1, product2}

	mockApp.On("GetAllProducts").Return(productsMock, nil)
	result, _ := mockApp.GetAllProducts()

	for r, val := range result {
		assert.Equal(t, val, productsExpected[r])
	}
	mockApp.AssertExpectations(t)
}

func TestGetProductByID(t *testing.T) {
	mockApp := mocks.NewProductAppInterface(t)
	mockApp.On("GetProductByID", mock.AnythingOfType("uuid.UUID")).Return(&product2, nil)
	testVal, err := mockApp.GetProductByID(product2.ProductId)
	if err == nil {
		assert.Equal(t, &product2, testVal)
	}
}

func TestAddProduct(t *testing.T) {
	mockApp := mocks.NewProductAppInterface(t)
	mockApp.On("AddProduct", mock.AnythingOfType("*entity.Product")).Return(&product2, nil)
	testVal, err := mockApp.AddProduct(&product2)
	if err == nil {
		assert.Equal(t, &product2, testVal)
	}
}
