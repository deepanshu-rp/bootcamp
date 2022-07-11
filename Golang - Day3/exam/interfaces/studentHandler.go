package interfaces

import (
	"exam/application"
	"exam/domain/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// defines the dependencies that will be used
type StudentDependency struct {
	stu application.StudentAppInterface
}

func NewStudentDependency(stu application.StudentAppInterface) *StudentDependency {
	return &StudentDependency{stu: stu}
}

func (s *StudentDependency) GetAllStudents(c *gin.Context) {
	students, err := s.stu.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (s *StudentDependency) AddStudent(c *gin.Context) {
	var student entity.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	// Validate input first

	if validationErr := student.Valdiate(); validationErr != nil {
		fmt.Println(validationErr)
		c.JSON(http.StatusNotAcceptable, gin.H{"error": validationErr.Error()})
		return
	}

	// Add to DB
	stu, err := s.stu.AddStudent(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, stu)
}

func (s *StudentDependency) GetStudentByID(c *gin.Context) {
	param := c.Params.ByName("id")

	fmt.Print("Param : ", param)
	id, e := uuid.FromString(param)

	// fmt.Print(id)
	if e != nil {
		c.AbortWithError(http.StatusNotFound, e)
	} else {
		if student, err := s.stu.GetStudentByID(id); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSONP(http.StatusOK, student)
		}
	}

}
