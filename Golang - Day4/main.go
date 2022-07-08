package main

import (
	"ecommerce/config"
	"ecommerce/models"
	"ecommerce/routes"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
)

func main() {

	// Connect to database

	orm, err := gorm.Open("mysql", config.DBUrl(config.DBSetup()))
	mu := &sync.Mutex{}

	if err != nil {
		fmt.Println(err)
	}
	config.DB = &config.Database{
		Orm: orm,
		Mu:  mu,
	}

	defer config.DB.Orm.Close()
	config.DB.Orm.AutoMigrate(&models.Product{})

	r := routes.SetUpRouter()
	//running
	r.Run()

}
