package models

import (
	"ecommerce/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetProductsDB(products *[]Product) {
	config.DB.Find(&products)
}
