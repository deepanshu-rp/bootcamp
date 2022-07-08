package models

import (
	"ecommerce/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetProductsDB(products *[]Product) {
	config.DB.Orm.Find(&products)
}

func AddProductDB(product *Product) error {
	err := config.DB.Orm.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}
