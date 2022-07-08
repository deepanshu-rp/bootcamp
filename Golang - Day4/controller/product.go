package controller

import (
	"ecommerce/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := models.GetProductsDB(&products); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSONP(http.StatusOK, gin.H{"products": products})
	}

}

func AddProduct(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	fmt.Println(product)

	if err := models.AddProductDB(&product); err != nil {
		c.JSONP(http.StatusNotFound, gin.H{
			"message": "not added"})
	} else {
		c.JSONP(http.StatusOK, gin.H{
			"id":           product.Id,
			"product_name": product.ProductName,
			"price":        product.Price,
			"quantity":     product.Quantity,
			"message":      "product added"})
	}

}
