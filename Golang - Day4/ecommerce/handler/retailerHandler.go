package handler

import (
	"ecommerce/application"
	"ecommerce/domain/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RetailerService struct {
	retail application.RetailerAppInterface
}

// args??
func NewRetailerService(rt application.RetailerAppInterface) *RetailerService {
	return &RetailerService{retail: rt}
}

func (rd *RetailerService) AddRetailer(c *gin.Context) {
	var retailer entity.Retailer
	// Bind request body
	if err := c.ShouldBindJSON(&retailer); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	// TODO: Validate retailer

	retailerNew, err := rd.retail.AddRetailer(&retailer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, retailerNew)
}

func (rd *RetailerService) GetRetailerByID(c *gin.Context) {
	param := c.Params.ByName("id")

	fmt.Print("Param : ", param)
	id, e := uuid.Parse(param)

	if e != nil {
		c.AbortWithError(http.StatusNotFound, e)
	} else {
		if retailer, err := rd.retail.GetRetailerByID(id); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSONP(http.StatusOK, retailer)
		}
	}

}
