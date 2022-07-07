package controllers

import (
	"exercise2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all students
func GetAllStudents(c *gin.Context) {
	var students []models.Student
	if err := models.GetAllStudents(&students); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSONP(http.StatusOK, students)

}

// Add new student
func AddNewStudent(c *gin.Context) {
	var student models.Student
	c.BindJSON(&student)
	err := models.AddNewStudent(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

// Get student by id
func GetStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	if err := models.GetStudent(&student, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSONP(http.StatusOK, student)
}

// Delete student entry
func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Request.URL.Query()["id"][0]
	err := models.DeleteStudent(id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	} else {
		c.JSONP(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

// Update single column
func UpdateMarks(c *gin.Context) {
	var student models.Student
	id := c.Request.URL.Query()["id"][0]
	marks := (c.Request.URL.Query()["marks"][0])
	println(marks)

	if e := models.UpdateMarks(id, marks); e != nil {
		c.JSON(http.StatusNotModified, student)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is updated"})
		println(marks)
	}

}

// Update all fields
func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	err := models.GetStudent(&student, id)
	// check if student exists
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	// update new value
	c.BindJSON(&student)
	err = models.UpdateStudent(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSONP(http.StatusOK, student)
	}
}
