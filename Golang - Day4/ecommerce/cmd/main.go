package main

import (
	"ecommerce/infrastructure/persistence"
	"ecommerce/interfaces"
	"errors"

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
	products := interfaces.NewProductDependency(service.ProductRepo, service.RetailerRepo)
	orders := interfaces.NewOrderDependency(service.OrderRepo)

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
		productsRoutes := gp.Group("/product")
		{
			productsRoutes.POST("/add", products.AddProduct)
			productsRoutes.GET("/all", products.GetAllProducts)
			productsRoutes.GET("/:id", products.GetProductByID)
		}
		orderRoutes := gp.Group("/order")
		{
			orderRoutes.GET("/:id", orders.GetOrderByID)
		}
	}

	err = router.Run()
	if err != nil {
		panic(errors.New("Unable to configure router"))
	}

}
