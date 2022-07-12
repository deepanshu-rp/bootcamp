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
	retailers := interfaces.NewRetailerDependency(service.RetailerRepo)

	router := gin.Default()
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"customer1": "pass",
	}))

	gp := authorized.Group("/")
	{
		customerRoutes := gp.Group("/customer")
		{
			customerRoutes.POST("/add", customers.AddCustomer)
			customerRoutes.GET("/:id", customers.GetCustomerByID)
		}
		retailerRoutes := gp.Group("/retailer")
		{
			retailerRoutes.POST("/add", retailers.AddRetailer)
			retailerRoutes.GET("/:id", retailers.GetRetailerByID)
		}
	}
	router.Run()

}
