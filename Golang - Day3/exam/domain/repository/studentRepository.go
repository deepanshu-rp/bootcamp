package repository

import (
	"errors"
	"exam/domain/entity"

	uuid "github.com/satori/go.uuid"
)

var (
	ErrStudentNotFound = errors.New("student was not found in the repository")
	ErrAddStudent      = errors.New("failed to add the student to the repository")
	ErrUpdateStudent   = errors.New("failed to update the student in the repository")
	ErrFetchFailed     = errors.New("could not fetch students from the repository")
)

type StudentRepository interface {
	GetStudentByID(uuid.UUID) (*entity.Student, error)
	AddStudent(*entity.Student) (*entity.Student, error)
	GetAllStudents() ([]entity.Student, error)
}
