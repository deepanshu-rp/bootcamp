package models

import (
	"ecommerce/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetProductsDB(products *[]Product) error {
	if err := config.DB.Orm.Find(&products).Error; err != nil {
		return err
	}
	return nil
}

func AddProductDB(product *Product) error {
	err := config.DB.Orm.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}
