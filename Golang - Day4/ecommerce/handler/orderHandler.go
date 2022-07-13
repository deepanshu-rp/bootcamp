package handler

import (
	"ecommerce/application"
	"ecommerce/domain/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderService struct {
	order    application.OrderAppInterface
	customer application.CustomerAppInterface
	product  application.ProductAppInterface
}

func NewOrderService(orderApp application.OrderAppInterface,
	customerApp application.CustomerAppInterface,
	productApp application.ProductAppInterface) *OrderService {
	return &OrderService{order: orderApp, customer: customerApp, product: productApp}
}

func (od *OrderService) AddMultipleOrders(c *gin.Context) {
	var orderRecieved entity.OrderRecieved
	var result []entity.OrderDetail

	// Bind request body
	if err := c.ShouldBindJSON(&orderRecieved); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// Validate customer id
	_, err := od.customer.GetCustomerByID(orderRecieved.CustomerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add to order table
	order, e := od.order.AddOrder(&entity.Order{OrderId: orderRecieved.OrderId, CustomerId: orderRecieved.CustomerId})
	if e != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": e.Error()})
		return
	}
	fmt.Println(orderRecieved)

	// Add order details
	for _, ord := range orderRecieved.Specs {
		// Check quantity
		fmt.Println(ord.ProductId)
		status := "not placed"
		product, err := od.product.GetProductByID(ord.ProductId)
		if err == nil && product.ProductQuantity-ord.ProductQuantity >= 0 {
			print(product.ProductQuantity - ord.ProductQuantity)
			product.ProductQuantity -= ord.ProductQuantity
			// Patch product
			if _, patchErr := od.product.UpdateProduct(product); patchErr == nil {
				status = "placed"
			}
		}
		ordPlaced, err := od.order.AddOrderDetail(&entity.OrderDetail{
			OrderId:         order.OrderId,
			OrderStatus:     status,
			ProductId:       ord.ProductId,
			QuantityOrdered: ord.ProductQuantity,
		})
		if err == nil {
			result = append(result, *ordPlaced)
		}

	}
	c.JSON(http.StatusOK, result)
}

func (od *OrderService) GetOrderByID(c *gin.Context) {
	param := c.Params.ByName("id")

	fmt.Print("Param : ", param)
	id, e := uuid.Parse(param)

	// Parse uuid recieved in param
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	} else {
		if order, err := od.order.GetOrderByID(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSONP(http.StatusOK, order)
		}
	}

}
