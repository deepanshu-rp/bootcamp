package routes

import (
	"exercise2/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	gp := router.Group("/")
	{
		gp.GET("student/:id", controllers.GetStudent)
		gp.GET("students", controllers.GetAllStudents)

		gp.POST("add", controllers.AddNewStudent)

		// Patch
		gp.PATCH("update-marks", controllers.UpdateMarks)
		gp.PUT("update-student/:id", controllers.UpdateStudent)

		gp.DELETE("delete", controllers.DeleteStudent)
	}
	return router
}
