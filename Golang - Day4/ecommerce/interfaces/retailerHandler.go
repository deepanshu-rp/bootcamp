package interfaces

import (
	"ecommerce/application"
	"ecommerce/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RetailerDependency struct {
	retail application.RetailerAppInterface
}

// args??
func NewRetailerDependency(rt application.RetailerAppInterface) *RetailerDependency {
	return &RetailerDependency{retail: rt}
}

func (rd *RetailerDependency) AddRetailer(c *gin.Context) {
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
