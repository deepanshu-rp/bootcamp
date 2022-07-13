package handler

import (
	"ecommerce/application"
	"ecommerce/domain/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomerService struct {
	cust application.CustomerAppInterface
}

func NewCustomerService(cd application.CustomerAppInterface) *CustomerService {
	return &CustomerService{cust: cd}
}

func (cd *CustomerService) AddCustomer(c *gin.Context) {
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

func (cd *CustomerService) GetCustomerByID(c *gin.Context) {
	param := c.Params.ByName("id")

	fmt.Print("Param : ", param)
	id, e := uuid.Parse(param)

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
