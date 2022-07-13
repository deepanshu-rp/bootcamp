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
	order1 = entity.Order{
		CustomerId: uuid.MustParse("e9ea26fd-12ab-44e7-baf8-77890d7d27ac"),
		OrderId:    uuid.MustParse("e3547962-01aa-4418-8b03-21103a9cedc1"),
	}
	order2 = entity.Order{
		OrderId:    uuid.MustParse("e3547962-01aa-4418-8b03-21103a9cedc1"),
		CustomerId: uuid.MustParse("e9ea26fd-12ab-44e7-baf8-77890d7d27ac"),
	}
)

var (
	orderDetail1 = entity.OrderDetail{
		OrderId:         uuid.MustParse("e3547962-01aa-4418-8b03-21103a9cedc1"),
		ProductId:       uuid.MustParse("e9ea26fd-12ab-44e7-baf8-77890d7d27ac"),
		QuantityOrdered: 10,
		OrderStatus:     "placed",
	}
	orderDetail2 = entity.OrderDetail{
		OrderId:         uuid.MustParse("e3547962-01aa-4418-8b03-21103a9cedc1"),
		ProductId:       uuid.MustParse("e9ea26fd-12ab-44e7-baf8-77890d7d27ac"),
		QuantityOrdered: 9,
		OrderStatus:     "not placed",
	}
)

func TestGetOrderByID(t *testing.T) {
	orderDetails := []entity.OrderDetail{orderDetail1, orderDetail2}
	mockApp := mocks.NewOrderAppInterface(t)
	mockApp.On("GetOrderByID", mock.AnythingOfType("uuid.UUID")).Return(orderDetails, nil)
	testVal, err := mockApp.GetOrderByID(uuid.New())
	if err == nil {
		assert.Equal(t, orderDetails, testVal)
	}
}

func TestAddOrder(t *testing.T) {
	mockApp := mocks.NewOrderAppInterface(t)
	mockApp.On("AddOrder", mock.AnythingOfType("*entity.Order")).Return(&order1, nil)
	testVal, err := mockApp.AddOrder(&entity.Order{})
	if err == nil {
		assert.Equal(t, &order1, testVal)
	}
}

func TestAddOrderDetail(t *testing.T) {
	mockApp := mocks.NewOrderAppInterface(t)
	mockApp.On("AddOrderDetail", mock.AnythingOfType("*entity.OrderDetail")).Return(&orderDetail2, nil)
	testVal, err := mockApp.AddOrderDetail(&entity.OrderDetail{})
	if err == nil {
		assert.Equal(t, &orderDetail2, testVal)
	}
}
