package main

import (
	"ecommerce/infrastructure/persistence"
	"ecommerce/interfaces"

	"github.com/gin-gonic/gin"
)

func main() {
	service, err := persistence.NewRepositories("mysql", "root", "password@123", "localhost", "ecommerce", 3306)
	if err != nil {
		panic(err)
	}
	defer service.Close()

	customers := interfaces.NewCustomerDependency(service.CustomerRepo)

	router := gin.Default()
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"customer1": "pass",
	}))

	gp := authorized.Group("/")
	{
		customerRoutes := gp.Group("/customer")
		{
			customerRoutes.POST("/add", customers.AddCustomer)
			customerRoutes.GET("/find/:id", customers.GetCustomerByID)
		}
	}
	router.Run()

}
