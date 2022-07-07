package main

import (
	"exercise2/config"
	"exercise2/models"
	"exercise2/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DBUrl(config.BuildDBConfig()))
	if err != nil {
		panic(any("Database error"))
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Student{})
	r := routes.SetUpRouter()
	r.Run()
}
