package models

type Product struct {
	Id          string `json:"id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
}
