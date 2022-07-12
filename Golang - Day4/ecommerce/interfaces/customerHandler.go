package interfaces

import (
	"ecommerce/application"
	"ecommerce/domain/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type CustomerDependency struct {
	cust application.CustomerAppInterface
}

func NewCustomerDependency(cd application.CustomerAppInterface) *CustomerDependency {
	return &CustomerDependency{cust: cd}
}

func (cd *CustomerDependency) AddCustomer(c *gin.Context) {
	var customer entity.Customer

	// Bind request body
	if err := c.ShouldBindJSON(&customer); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	// TODO: Validate customer

	// TODO: DB ops
	cust, err := cd.cust.AddCustomer(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, cust)

}

func (cd *CustomerDependency) GetCustomerByID(c *gin.Context) {
	param := c.Params.ByName("id")

	fmt.Print("Param : ", param)
	id, e := uuid.FromString(param)

	if e != nil {
		c.AbortWithError(http.StatusNotFound, e)
	} else {
		if customer, err := cd.cust.GetCustomerByID(id); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSONP(http.StatusOK, customer)
		}
	}

}
