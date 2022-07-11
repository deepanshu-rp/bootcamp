package application

import (
	"exam/domain/entity"
	"exam/domain/repository"
)

type marksApp struct {
	mk repository.MarksRepository
}

type MarksAppInterface interface {
	AddMarks(*entity.Marks) (*entity.Marks, error)
}

var _ MarksAppInterface = &marksApp{}

func (m *marksApp) AddMarks(marks *entity.Marks) (*entity.Marks, error) {
	return m.mk.AddMarks(marks)
}

// TODO: delete marks by stuednt ID
