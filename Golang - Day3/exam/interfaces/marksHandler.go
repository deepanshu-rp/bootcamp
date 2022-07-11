package interfaces

import (
	"exam/application"
	"exam/domain/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MarksDependency struct {
	stu application.StudentAppInterface
	mks application.MarksAppInterface
}

func NewMarksDependency(stu application.StudentAppInterface, mks application.MarksAppInterface) *MarksDependency {
	return &MarksDependency{stu: stu, mks: mks}
}

func (m *MarksDependency) AddMarks(c *gin.Context) {
	var mark entity.Marks

	if err := c.ShouldBindJSON(&mark); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	// Validate input
	if validationErr := mark.Valdiate(); validationErr != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": validationErr.Error()})
	}

	// Validate student
	id := mark.StudentId
	_, err := m.stu.GetStudentByID(id)

	if err != nil {
		// Student doesn't exist
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Add to DB
	marks, err := m.mks.AddMarks(&mark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, marks)
}
