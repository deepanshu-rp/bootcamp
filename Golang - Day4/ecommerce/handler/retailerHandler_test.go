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
	retailer1 = entity.Retailer{
		RetailerId:     uuid.MustParse("e3547962-01aa-4418-8b03-21103a9cedc1"),
		RetailerName:   "Test 1",
		RetailerRating: 4,
	}
	retailer2 = entity.Retailer{
		RetailerId:     uuid.MustParse("e9ea26fd-12ab-44e7-baf8-77890d7d27ac"),
		RetailerName:   "Test 2",
		RetailerRating: 3,
	}
)

func TestGetRetailerByID(t *testing.T) {
	mockApp := mocks.NewRetailerAppInterface(t)
	mockApp.On("GetRetailerByID", mock.AnythingOfType("uuid.UUID")).Return(&retailer1, nil)
	testApp := handler.NewRetailerService(mockApp)

	retailers := []entity.Retailer{retailer1, retailer1}
	for _, r := range retailers {
		result, _ := testApp.Retail.GetRetailerByID(r.RetailerId)
		assert.Equal(t, result, &r)
	}
}

func TestAddReatiler(t *testing.T) {
	mockApp := mocks.NewRetailerAppInterface(t)
	mockApp.On("AddRetailer", mock.AnythingOfType("*entity.Retailer")).Return(
		func(r *entity.Retailer) *entity.Retailer {
			return r
		},
		func(r *entity.Retailer) error {
			return nil
		},
	)
	retailersTest := []entity.Retailer{retailer1, retailer2}

	testApp := handler.NewRetailerService(mockApp)
	for _, r := range retailersTest {
		result, _ := testApp.Retail.AddRetailer(&r)
		assert.Equal(t, result, &r)
	}
}
