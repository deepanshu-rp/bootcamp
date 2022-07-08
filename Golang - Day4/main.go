package main

import (
	"ecommerce/config"

	"github.com/jinzhu/gorm"
)

func main() {
	var err error

	// Connect to database
	config.DB, err = gorm.Open("mysql", config.DBUrl(config.DBSetup()))

	if err != nil {
		panic(any("Database error"))
	}

	defer config.DB.Close()

}
