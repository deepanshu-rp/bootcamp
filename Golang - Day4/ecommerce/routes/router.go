package routes

import (
	"ecommerce/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(
	customerService *handler.CustomerService,
	retailerService *handler.RetailerService,
	productService *handler.ProductService,
	orderService *handler.OrderService) *gin.Engine {

	router := gin.Default()
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"customer1": "pass",
	}))

	gp := authorized.Group("/")
	{
		customerRoutes := gp.Group("/customer")
		{
			customerRoutes.POST("/add", customerService.AddCustomer)
			customerRoutes.GET("/:id", customerService.GetCustomerByID)
		}
		retailerRoutes := gp.Group("/retailer")
		{
			retailerRoutes.POST("/add", retailerService.AddRetailer)
			retailerRoutes.GET("/:id", retailerService.GetRetailerByID)
		}
		productsRoutes := gp.Group("/product")
		{
			productsRoutes.POST("/add", productService.AddProduct)
			productsRoutes.GET("/all", productService.GetAllProducts)
			productsRoutes.GET("/:id", productService.GetProductByID)
			productsRoutes.PATCH("/:id", productService.UpdateProduct)
		}
		orderRoutes := gp.Group("/order")
		{
			orderRoutes.POST("/add", orderService.AddMultipleOrders)
			orderRoutes.GET("/:id", orderService.GetOrderByID)
		}
	}
	return router
}
