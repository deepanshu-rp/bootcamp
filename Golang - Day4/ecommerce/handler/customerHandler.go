package handler

import (
	"ecommerce/application"
	"ecommerce/domain/entity"
	"ecommerce/handler/utils"
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
		c.JSON(http.StatusBadRequest, utils.FormatErrorResponse(entity.ErrInvalidCustomerJSON.Error(), err.Error()))
		return
	}

	// Validate customer
	if customErr, err := customer.ValidateInput(); err != nil {
		c.JSON(http.StatusBadRequest, utils.FormatErrorResponse(customErr.Error(), err.Error()))
		return
	}

	// Add to DB
	cust, err := cd.cust.AddCustomer(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.FormatErrorResponse(entity.ErrCustomerNotAdded.Error(), err.Error()))
		return
	}
	c.JSON(http.StatusCreated, cust)

}

func (cd *CustomerService) GetCustomerByID(c *gin.Context) {
	param := c.Params.ByName("id")

	fmt.Print("Param : ", param)
	id, e := uuid.Parse(param)

	if e != nil {
		c.JSON(http.StatusBadRequest, utils.FormatErrorResponse(entity.ErrInvalidCustomerUUID.Error(), e.Error()))
	} else {
		if customer, err := cd.cust.GetCustomerByID(id); err != nil {
			c.JSON(http.StatusNotFound, utils.FormatErrorResponse(entity.ErrInvalidCustomerUUID.Error(), err.Error()))
		} else {
			c.JSONP(http.StatusOK, customer)
		}
	}

}
