package interfaces

import (
	"ecommerce/application"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderDependency struct {
	order application.OrderAppInterface
}

func NewOrderDependency(orderApp application.OrderAppInterface) *OrderDependency {
	return &OrderDependency{order: orderApp}
}

func (od *OrderDependency) GetOrderByID(c *gin.Context) {
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
