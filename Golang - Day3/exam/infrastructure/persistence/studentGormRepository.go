package persistence

import (
	"exam/domain/entity"
	"exam/domain/repository"

	uuid "github.com/satori/go.uuid"

	"github.com/jinzhu/gorm"
)

type StudentGormRepo struct {
	db *gorm.DB
}

func NewStudentGormRepository(db *gorm.DB) *StudentGormRepo {
	return &StudentGormRepo{db}
}

var _ repository.StudentRepository = &StudentGormRepo{}

func (s *StudentGormRepo) GetAllStudents() ([]entity.Student, error) {
	var students []entity.Student
	err := s.db.Debug().Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudentGormRepo) AddStudent(student *entity.Student) (*entity.Student, error) {
	err := s.db.Create(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentGormRepo) GetStudentByID(uuid uuid.UUID) (*entity.Student, error) {
	var student entity.Student
	err := s.db.Where("id = ?", uuid).Find(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}
