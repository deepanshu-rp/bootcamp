package controller

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	models.GetProductsDB(&products)
	c.JSONP(http.StatusOK, gin.H{"products": products})
}
