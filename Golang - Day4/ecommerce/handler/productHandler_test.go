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
	mockProductApp := mocks.NewProductAppInterface(t)
	mockRetailerApp := mocks.NewRetailerAppInterface(t)

	mockProductApp.On("UpdateProduct", mock.AnythingOfType("*entity.Product")).Return(&product1, nil)

	testApp := handler.NewProductService(mockProductApp, mockRetailerApp)
	testVal, _ := testApp.Product.UpdateProduct(&product2)
	assert.Equal(t, &product1, testVal)
}

func TestGetAllProducts(t *testing.T) {
	mockProductApp := mocks.NewProductAppInterface(t)
	mockRetailerApp := mocks.NewRetailerAppInterface(t)

	productsMock := []entity.Product{product1, product2}
	productsExpected := []entity.Product{product1, product2}

	mockProductApp.On("GetAllProducts").Return(productsMock, nil)
	testApp := handler.NewProductService(mockProductApp, mockRetailerApp)

	result, _ := testApp.Product.GetAllProducts()

	for r, val := range result {
		assert.Equal(t, val, productsExpected[r])
	}
}

func TestGetProductByID(t *testing.T) {
	mockProductApp := mocks.NewProductAppInterface(t)
	mockRetailerApp := mocks.NewRetailerAppInterface(t)

	mockProductApp.On("GetProductByID", mock.AnythingOfType("uuid.UUID")).Return(&product1, nil)

	testApp := handler.NewProductService(mockProductApp, mockRetailerApp)

	productsTest := []entity.Product{product1, product1}
	for _, p := range productsTest {
		result, _ := testApp.Product.GetProductByID(p.ProductId)
		assert.Equal(t, result, &product1)
	}
}

func TestAddProduct(t *testing.T) {
	mockProductApp := mocks.NewProductAppInterface(t)
	mockRetailerApp := mocks.NewRetailerAppInterface(t)

	mockProductApp.On("AddProduct", mock.AnythingOfType("*entity.Product")).Return(
		func(p *entity.Product) *entity.Product {
			return p
		},
		func(p *entity.Product) error {
			return nil
		},
	)

	testApp := handler.NewProductService(mockProductApp, mockRetailerApp)

	productsTest := []entity.Product{product1, product2}

	for _, p := range productsTest {
		result, _ := testApp.Product.AddProduct(&p)
		assert.Equal(t, result, &p)
	}
}
