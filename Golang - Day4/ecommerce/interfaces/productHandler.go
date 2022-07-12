package interfaces

import (
	"ecommerce/application"
	"ecommerce/domain/entity"
	"ecommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductDependency struct {
	product application.ProductAppInterface
	retail  application.RetailerAppInterface
}

func NewProductDependency(pdt application.ProductAppInterface, rt application.RetailerAppInterface) *ProductDependency {
	return &ProductDependency{product: pdt, retail: rt}
}

func (pd *ProductDependency) AddProduct(c *gin.Context) {
	var product entity.Product
	// Bind request body
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	// TODO: Validate product

	// Check if retailer valid

	if _, err := pd.retail.GetRetailerByID(uuid.UUID(product.RetailerId)); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	productNew, err := pd.product.AddProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	response := utils.FormatProductAddedResponse(*productNew)
	c.JSON(http.StatusCreated, response)
}

func (pd *ProductDependency) GetAllProducts(c *gin.Context) {

	products, err := pd.product.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := utils.FormatProductListResponse(products)
	c.JSONP(http.StatusOK, gin.H{"products": result})
}
