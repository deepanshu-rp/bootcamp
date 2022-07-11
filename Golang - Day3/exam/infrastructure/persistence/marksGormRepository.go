package persistence

import (
	"exam/domain/entity"
	"exam/domain/repository"

	"github.com/jinzhu/gorm"
)

type MarksGormRepo struct {
	db *gorm.DB
}

func NewMarksGormRepository(db *gorm.DB) *MarksGormRepo {
	return &MarksGormRepo{db}
}

var _ repository.MarksRepository = &MarksGormRepo{}

func (m *MarksGormRepo) AddMarks(marks *entity.Marks) (*entity.Marks, error) {

	if err := m.db.Create(&marks).Error; err != nil {
		return nil, err
	}
	return marks, nil

}
