package application

import (
	"exam/domain/entity"
	"exam/domain/repository"

	uuid "github.com/satori/go.uuid"
)

type studentApp struct {
	stu repository.StudentRepository
}

type StudentAppInterface interface {
	GetAllStudents() ([]entity.Student, error)
	AddStudent(*entity.Student) (*entity.Student, error)
	GetStudentByID(uuid uuid.UUID) (*entity.Student, error)
}

var _ StudentAppInterface = &studentApp{}

func (s *studentApp) GetAllStudents() ([]entity.Student, error) {
	// get student from repository
	return s.stu.GetAllStudents()
}

func (s *studentApp) AddStudent(student *entity.Student) (*entity.Student, error) {
	return s.stu.AddStudent(student)
}

func (s *studentApp) GetStudentByID(uuid uuid.UUID) (*entity.Student, error) {
	return s.stu.GetStudentByID(uuid)
}

// TODO: Delete student
