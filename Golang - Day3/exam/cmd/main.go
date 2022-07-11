package main

import (
	"exam/infrastructure/persistence"
	"exam/interfaces"

	"github.com/gin-gonic/gin"
)

func main() {
	service, err := persistence.NewRepositories("mysql", "root", "password@123", "localhost", "exams", 3306)
	if err != nil {
		panic(err)
	}
	defer service.Close()
	service.Automigrate()

	students := interfaces.NewStudentDependency(service.Student)
	marks := interfaces.NewMarksDependency(service.Student, service.Marks)
	r := gin.Default()
	gpStudent := r.Group("/student")
	{
		gpStudent.GET("/all", students.GetAllStudents)
		gpStudent.POST("/add", students.AddStudent)
		gpStudent.GET("/find/:id", students.GetStudentByID)
	}
	gpMarks := r.Group("/marks")
	{
		gpMarks.POST("/add", marks.AddMarks)
	}

	r.Run()

}
