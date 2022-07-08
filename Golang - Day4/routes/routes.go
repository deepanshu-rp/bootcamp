package routes

import (
	"ecommerce/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	group := router.Group("/")
	{
		group.GET("products", controller.GetProducts)
		group.POST("products/add", controller.AddProduct)
	}
	return router

}
