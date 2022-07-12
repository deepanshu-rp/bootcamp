package utils

import (
	"ecommerce/domain/entity"

	"github.com/google/uuid"
)

type ProductListResponse struct {
	ProductId       uuid.UUID `json:"product_id"`
	ProductName     string    `json:"product_name"`
	ProductPrice    int       `json:"product_price"`
	ProductQuantity int       `json:"product_quantity"`
}
type ProductAddResponse struct {
	ProductId       uuid.UUID `json:"product_id"`
	ProductName     string    `json:"product_name"`
	ProductPrice    int       `json:"product_price"`
	ProductQuantity int       `json:"product_quantity"`
	Message         string    `json:"message"`
}

func FormatProductListResponse(products []entity.Product) []ProductListResponse {
	var response []ProductListResponse
	for _, p := range products {
		val := ProductListResponse{
			ProductId:       p.ProductId,
			ProductName:     p.ProductName,
			ProductPrice:    p.ProductPrice,
			ProductQuantity: p.ProductQuantity,
		}
		response = append(response, val)
	}
	return response
}

func FormatProductAddedResponse(product entity.Product) ProductAddResponse {
	return ProductAddResponse{
		ProductId:       product.ProductId,
		ProductName:     product.ProductName,
		ProductPrice:    product.ProductPrice,
		ProductQuantity: product.ProductQuantity,
		Message:         "added successfully",
	}

}
