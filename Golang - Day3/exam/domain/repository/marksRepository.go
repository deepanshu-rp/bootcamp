package repository

import "exam/domain/entity"

type MarksRepository interface {
	AddMarks(*entity.Marks) (*entity.Marks, error)
}
