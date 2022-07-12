package interfaces

import (
	"ecommerce/application"
	"ecommerce/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type productsResponse struct {
	ProductId       uuid.UUID `json:"product_id"`
	ProductName     string    `json:"product_name"`
	ProductPrice    int       `json:"product_price"`
	ProductQuantity int       `json:"product_quantity"`
}

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
	c.JSON(http.StatusCreated, gin.H{
		"id":           productNew.ProductId,
		"product_name": productNew.ProductName,
		"price":        productNew.ProductPrice,
		"quantity":     productNew.ProductQuantity,
		"message":      "product successfully added",
	})
}

func (pd *ProductDependency) GetAllProducts(c *gin.Context) {

	products, err := pd.product.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var result []productsResponse
	for _, p := range products {
		val := productsResponse{
			ProductId:       p.ProductId,
			ProductName:     p.ProductName,
			ProductPrice:    p.ProductPrice,
			ProductQuantity: p.ProductQuantity,
		}
		result = append(result, val)
	}
	c.JSONP(http.StatusOK, gin.H{"products": result})
}
