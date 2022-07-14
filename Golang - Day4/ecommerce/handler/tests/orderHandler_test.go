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
	mockOrderApp := mocks.NewOrderAppInterface(t)
	mockCustomerApp := mocks.NewCustomerAppInterface(t)
	mockProductApp := mocks.NewProductAppInterface(t)
	testApp := handler.NewOrderService(mockOrderApp, mockCustomerApp, mockProductApp)

	mockOrderApp.On("GetOrderByID", mock.AnythingOfType("uuid.UUID")).Return(orderDetails, nil)

	testVal, _ := testApp.Order.GetOrderByID(uuid.New())
	assert.Equal(t, orderDetails, testVal)
}

func TestAddOrder(t *testing.T) {
	mockOrderApp := mocks.NewOrderAppInterface(t)
	mockCustomerApp := mocks.NewCustomerAppInterface(t)
	mockProductApp := mocks.NewProductAppInterface(t)

	mockOrderApp.On("AddOrder", mock.AnythingOfType("*entity.Order")).Return(
		func(o *entity.Order) *entity.Order {
			return o
		},
		func(o *entity.Order) error {
			return nil
		},
	)

	testApp := handler.NewOrderService(mockOrderApp, mockCustomerApp, mockProductApp)
	orderTests := []entity.Order{order1, order2}

	for _, o := range orderTests {
		testVal, _ := testApp.Order.AddOrder(&o)
		assert.Equal(t, &o, testVal)
	}
}

func TestAddOrderDetail(t *testing.T) {
	mockOrderApp := mocks.NewOrderAppInterface(t)
	mockCustomerApp := mocks.NewCustomerAppInterface(t)
	mockProductApp := mocks.NewProductAppInterface(t)

	mockOrderApp.On("AddOrderDetail", mock.AnythingOfType("*entity.OrderDetail")).Return(
		func(od *entity.OrderDetail) *entity.OrderDetail {
			return od
		},
		func(o *entity.OrderDetail) error {
			return nil
		},
	)

	testApp := handler.NewOrderService(mockOrderApp, mockCustomerApp, mockProductApp)
	orderDetailTests := []entity.OrderDetail{orderDetail1, orderDetail2}

	for _, od := range orderDetailTests {
		testVal, _ := testApp.Order.AddOrderDetail(&od)
		assert.Equal(t, &od, testVal)
	}

}
